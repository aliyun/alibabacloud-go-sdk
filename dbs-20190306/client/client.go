// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConfigureBackupPlanRequest struct {
	BackupPlanId                      *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	SourceEndpointInstanceType        *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	SourceEndpointRegion              *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	SourceEndpointInstanceID          *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	SourceEndpointIP                  *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	SourceEndpointPort                *int32  `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	SourceEndpointDatabaseName        *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	SourceEndpointUserName            *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	SourceEndpointPassword            *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	BackupGatewayId                   *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	OSSBucketName                     *string `json:"OSSBucketName,omitempty" xml:"OSSBucketName,omitempty"`
	BackupObjects                     *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	BackupPeriod                      *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupStartTime                   *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	EnableBackupLog                   *bool   `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	BackupLogIntervalSeconds          *int32  `json:"BackupLogIntervalSeconds,omitempty" xml:"BackupLogIntervalSeconds,omitempty"`
	BackupRetentionPeriod             *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	DuplicationInfrequentAccessPeriod *int32  `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	DuplicationArchivePeriod          *int32  `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	BackupPlanName                    *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	SourceEndpointOracleSID           *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	BackupStorageType                 *string `json:"BackupStorageType,omitempty" xml:"BackupStorageType,omitempty"`
	BackupSpeedLimit                  *int64  `json:"BackupSpeedLimit,omitempty" xml:"BackupSpeedLimit,omitempty"`
	BackupRateLimit                   *int64  `json:"BackupRateLimit,omitempty" xml:"BackupRateLimit,omitempty"`
	BackupStrategyType                *string `json:"BackupStrategyType,omitempty" xml:"BackupStrategyType,omitempty"`
	CrossAliyunId                     *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	CrossRoleName                     *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	AutoStartBackup                   *bool   `json:"AutoStartBackup,omitempty" xml:"AutoStartBackup,omitempty"`
	OwnerId                           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken                       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ConfigureBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *ConfigureBackupPlanRequest) SetBackupPlanId(v string) *ConfigureBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointInstanceType(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointRegion(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointInstanceID(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointIP(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointPort(v int32) *ConfigureBackupPlanRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointDatabaseName(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointUserName(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointPassword(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupGatewayId(v int64) *ConfigureBackupPlanRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetOSSBucketName(v string) *ConfigureBackupPlanRequest {
	s.OSSBucketName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupObjects(v string) *ConfigureBackupPlanRequest {
	s.BackupObjects = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupPeriod(v string) *ConfigureBackupPlanRequest {
	s.BackupPeriod = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupStartTime(v string) *ConfigureBackupPlanRequest {
	s.BackupStartTime = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetEnableBackupLog(v bool) *ConfigureBackupPlanRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupLogIntervalSeconds(v int32) *ConfigureBackupPlanRequest {
	s.BackupLogIntervalSeconds = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupRetentionPeriod(v int32) *ConfigureBackupPlanRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetDuplicationInfrequentAccessPeriod(v int32) *ConfigureBackupPlanRequest {
	s.DuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetDuplicationArchivePeriod(v int32) *ConfigureBackupPlanRequest {
	s.DuplicationArchivePeriod = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupPlanName(v string) *ConfigureBackupPlanRequest {
	s.BackupPlanName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointOracleSID(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupStorageType(v string) *ConfigureBackupPlanRequest {
	s.BackupStorageType = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupSpeedLimit(v int64) *ConfigureBackupPlanRequest {
	s.BackupSpeedLimit = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupRateLimit(v int64) *ConfigureBackupPlanRequest {
	s.BackupRateLimit = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupStrategyType(v string) *ConfigureBackupPlanRequest {
	s.BackupStrategyType = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetCrossAliyunId(v string) *ConfigureBackupPlanRequest {
	s.CrossAliyunId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetCrossRoleName(v string) *ConfigureBackupPlanRequest {
	s.CrossRoleName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetAutoStartBackup(v bool) *ConfigureBackupPlanRequest {
	s.AutoStartBackup = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetOwnerId(v string) *ConfigureBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetClientToken(v string) *ConfigureBackupPlanRequest {
	s.ClientToken = &v
	return s
}

type ConfigureBackupPlanResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ConfigureBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureBackupPlanResponseBody) SetRequestId(v string) *ConfigureBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetHttpStatusCode(v int32) *ConfigureBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetBackupPlanId(v string) *ConfigureBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetErrMessage(v string) *ConfigureBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetSuccess(v bool) *ConfigureBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureBackupPlanResponseBody) SetErrCode(v string) *ConfigureBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

type ConfigureBackupPlanResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *ConfigureBackupPlanResponse) SetHeaders(v map[string]*string) *ConfigureBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *ConfigureBackupPlanResponse) SetBody(v *ConfigureBackupPlanResponseBody) *ConfigureBackupPlanResponse {
	s.Body = v
	return s
}

type CreateBackupPlanRequest struct {
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceClass  *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	DatabaseType   *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	BackupMethod   *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	Period         *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime       *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	PayType        *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DatabaseRegion *string `json:"DatabaseRegion,omitempty" xml:"DatabaseRegion,omitempty"`
	StorageRegion  *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	StorageType    *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	FromApp        *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequest) SetRegion(v string) *CreateBackupPlanRequest {
	s.Region = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInstanceClass(v string) *CreateBackupPlanRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDatabaseType(v string) *CreateBackupPlanRequest {
	s.DatabaseType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetBackupMethod(v string) *CreateBackupPlanRequest {
	s.BackupMethod = &v
	return s
}

func (s *CreateBackupPlanRequest) SetPeriod(v string) *CreateBackupPlanRequest {
	s.Period = &v
	return s
}

func (s *CreateBackupPlanRequest) SetUsedTime(v int32) *CreateBackupPlanRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateBackupPlanRequest) SetPayType(v string) *CreateBackupPlanRequest {
	s.PayType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDatabaseRegion(v string) *CreateBackupPlanRequest {
	s.DatabaseRegion = &v
	return s
}

func (s *CreateBackupPlanRequest) SetStorageRegion(v string) *CreateBackupPlanRequest {
	s.StorageRegion = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInstanceType(v string) *CreateBackupPlanRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetStorageType(v string) *CreateBackupPlanRequest {
	s.StorageType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetFromApp(v string) *CreateBackupPlanRequest {
	s.FromApp = &v
	return s
}

func (s *CreateBackupPlanRequest) SetOwnerId(v string) *CreateBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetClientToken(v string) *CreateBackupPlanRequest {
	s.ClientToken = &v
	return s
}

type CreateBackupPlanResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s CreateBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponseBody) SetRequestId(v string) *CreateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetHttpStatusCode(v int32) *CreateBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetBackupPlanId(v string) *CreateBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetErrMessage(v string) *CreateBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetOrderId(v string) *CreateBackupPlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetSuccess(v bool) *CreateBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBackupPlanResponseBody) SetErrCode(v string) *CreateBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

type CreateBackupPlanResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponse) SetHeaders(v map[string]*string) *CreateBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupPlanResponse) SetBody(v *CreateBackupPlanResponseBody) *CreateBackupPlanResponse {
	s.Body = v
	return s
}

type CreateFullBackupSetDownloadRequest struct {
	BackupSetId         *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSetDataFormat *string `json:"BackupSetDataFormat,omitempty" xml:"BackupSetDataFormat,omitempty"`
	OwnerId             *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateFullBackupSetDownloadRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFullBackupSetDownloadRequest) GoString() string {
	return s.String()
}

func (s *CreateFullBackupSetDownloadRequest) SetBackupSetId(v string) *CreateFullBackupSetDownloadRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateFullBackupSetDownloadRequest) SetBackupSetDataFormat(v string) *CreateFullBackupSetDownloadRequest {
	s.BackupSetDataFormat = &v
	return s
}

func (s *CreateFullBackupSetDownloadRequest) SetOwnerId(v string) *CreateFullBackupSetDownloadRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFullBackupSetDownloadRequest) SetClientToken(v string) *CreateFullBackupSetDownloadRequest {
	s.ClientToken = &v
	return s
}

type CreateFullBackupSetDownloadResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupSetDownloadTaskId *string `json:"BackupSetDownloadTaskId,omitempty" xml:"BackupSetDownloadTaskId,omitempty"`
	HttpStatusCode          *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage              *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success                 *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode                 *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s CreateFullBackupSetDownloadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFullBackupSetDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFullBackupSetDownloadResponseBody) SetRequestId(v string) *CreateFullBackupSetDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetBackupSetDownloadTaskId(v string) *CreateFullBackupSetDownloadResponseBody {
	s.BackupSetDownloadTaskId = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetHttpStatusCode(v int32) *CreateFullBackupSetDownloadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetErrMessage(v string) *CreateFullBackupSetDownloadResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetSuccess(v bool) *CreateFullBackupSetDownloadResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFullBackupSetDownloadResponseBody) SetErrCode(v string) *CreateFullBackupSetDownloadResponseBody {
	s.ErrCode = &v
	return s
}

type CreateFullBackupSetDownloadResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFullBackupSetDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFullBackupSetDownloadResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFullBackupSetDownloadResponse) GoString() string {
	return s.String()
}

func (s *CreateFullBackupSetDownloadResponse) SetHeaders(v map[string]*string) *CreateFullBackupSetDownloadResponse {
	s.Headers = v
	return s
}

func (s *CreateFullBackupSetDownloadResponse) SetBody(v *CreateFullBackupSetDownloadResponseBody) *CreateFullBackupSetDownloadResponse {
	s.Body = v
	return s
}

type CreateGetDBListFromAgentTaskRequest struct {
	BackupGatewayId      *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	DatabaseType         *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	SourceEndpointIP     *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	SourceEndpointPort   *int32  `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateGetDBListFromAgentTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGetDBListFromAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateGetDBListFromAgentTaskRequest) SetBackupGatewayId(v int64) *CreateGetDBListFromAgentTaskRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetDatabaseType(v string) *CreateGetDBListFromAgentTaskRequest {
	s.DatabaseType = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetSourceEndpointRegion(v string) *CreateGetDBListFromAgentTaskRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetSourceEndpointIP(v string) *CreateGetDBListFromAgentTaskRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetSourceEndpointPort(v int32) *CreateGetDBListFromAgentTaskRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetOwnerId(v string) *CreateGetDBListFromAgentTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskRequest) SetClientToken(v string) *CreateGetDBListFromAgentTaskRequest {
	s.ClientToken = &v
	return s
}

type CreateGetDBListFromAgentTaskResponseBody struct {
	TaskId         *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s CreateGetDBListFromAgentTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGetDBListFromAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetTaskId(v int64) *CreateGetDBListFromAgentTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetRequestId(v string) *CreateGetDBListFromAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetHttpStatusCode(v int32) *CreateGetDBListFromAgentTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetErrMessage(v string) *CreateGetDBListFromAgentTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetSuccess(v bool) *CreateGetDBListFromAgentTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponseBody) SetErrCode(v string) *CreateGetDBListFromAgentTaskResponseBody {
	s.ErrCode = &v
	return s
}

type CreateGetDBListFromAgentTaskResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGetDBListFromAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGetDBListFromAgentTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGetDBListFromAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateGetDBListFromAgentTaskResponse) SetHeaders(v map[string]*string) *CreateGetDBListFromAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateGetDBListFromAgentTaskResponse) SetBody(v *CreateGetDBListFromAgentTaskResponseBody) *CreateGetDBListFromAgentTaskResponse {
	s.Body = v
	return s
}

type CreateIncrementBackupSetDownloadRequest struct {
	BackupSetId         *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSetName       *string `json:"BackupSetName,omitempty" xml:"BackupSetName,omitempty"`
	BackupSetDataFormat *string `json:"BackupSetDataFormat,omitempty" xml:"BackupSetDataFormat,omitempty"`
	OwnerId             *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateIncrementBackupSetDownloadRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIncrementBackupSetDownloadRequest) GoString() string {
	return s.String()
}

func (s *CreateIncrementBackupSetDownloadRequest) SetBackupSetId(v string) *CreateIncrementBackupSetDownloadRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) SetBackupSetName(v string) *CreateIncrementBackupSetDownloadRequest {
	s.BackupSetName = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) SetBackupSetDataFormat(v string) *CreateIncrementBackupSetDownloadRequest {
	s.BackupSetDataFormat = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) SetOwnerId(v string) *CreateIncrementBackupSetDownloadRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadRequest) SetClientToken(v string) *CreateIncrementBackupSetDownloadRequest {
	s.ClientToken = &v
	return s
}

type CreateIncrementBackupSetDownloadResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupSetDownloadTaskId *string `json:"BackupSetDownloadTaskId,omitempty" xml:"BackupSetDownloadTaskId,omitempty"`
	HttpStatusCode          *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage              *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success                 *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode                 *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s CreateIncrementBackupSetDownloadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIncrementBackupSetDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetRequestId(v string) *CreateIncrementBackupSetDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetBackupSetDownloadTaskId(v string) *CreateIncrementBackupSetDownloadResponseBody {
	s.BackupSetDownloadTaskId = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetHttpStatusCode(v int32) *CreateIncrementBackupSetDownloadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetErrMessage(v string) *CreateIncrementBackupSetDownloadResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetSuccess(v bool) *CreateIncrementBackupSetDownloadResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponseBody) SetErrCode(v string) *CreateIncrementBackupSetDownloadResponseBody {
	s.ErrCode = &v
	return s
}

type CreateIncrementBackupSetDownloadResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIncrementBackupSetDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIncrementBackupSetDownloadResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIncrementBackupSetDownloadResponse) GoString() string {
	return s.String()
}

func (s *CreateIncrementBackupSetDownloadResponse) SetHeaders(v map[string]*string) *CreateIncrementBackupSetDownloadResponse {
	s.Headers = v
	return s
}

func (s *CreateIncrementBackupSetDownloadResponse) SetBody(v *CreateIncrementBackupSetDownloadResponseBody) *CreateIncrementBackupSetDownloadResponse {
	s.Body = v
	return s
}

type CreateRestoreTaskRequest struct {
	BackupPlanId                    *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	DestinationEndpointRegion       *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	DestinationEndpointInstanceID   *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	DestinationEndpointIP           *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	DestinationEndpointPort         *int32  `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	DestinationEndpointDatabaseName *string `json:"DestinationEndpointDatabaseName,omitempty" xml:"DestinationEndpointDatabaseName,omitempty"`
	DestinationEndpointUserName     *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	DestinationEndpointPassword     *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	BackupGatewayId                 *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	RestoreObjects                  *string `json:"RestoreObjects,omitempty" xml:"RestoreObjects,omitempty"`
	RestoreTime                     *int64  `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	BackupSetId                     *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	RestoreTaskName                 *string `json:"RestoreTaskName,omitempty" xml:"RestoreTaskName,omitempty"`
	DestinationEndpointOracleSID    *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	RestoreDir                      *string `json:"RestoreDir,omitempty" xml:"RestoreDir,omitempty"`
	RestoreHome                     *string `json:"RestoreHome,omitempty" xml:"RestoreHome,omitempty"`
	DuplicateConflict               *string `json:"DuplicateConflict,omitempty" xml:"DuplicateConflict,omitempty"`
	CrossAliyunId                   *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	CrossRoleName                   *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	OwnerId                         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken                     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateRestoreTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRestoreTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRestoreTaskRequest) SetBackupPlanId(v string) *CreateRestoreTaskRequest {
	s.BackupPlanId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointInstanceType(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointRegion(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointInstanceID(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointIP(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointIP = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointPort(v int32) *CreateRestoreTaskRequest {
	s.DestinationEndpointPort = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointDatabaseName(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointDatabaseName = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointUserName(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointPassword(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointPassword = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetBackupGatewayId(v int64) *CreateRestoreTaskRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreObjects(v string) *CreateRestoreTaskRequest {
	s.RestoreObjects = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreTime(v int64) *CreateRestoreTaskRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetBackupSetId(v string) *CreateRestoreTaskRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreTaskName(v string) *CreateRestoreTaskRequest {
	s.RestoreTaskName = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointOracleSID(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreDir(v string) *CreateRestoreTaskRequest {
	s.RestoreDir = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreHome(v string) *CreateRestoreTaskRequest {
	s.RestoreHome = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDuplicateConflict(v string) *CreateRestoreTaskRequest {
	s.DuplicateConflict = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetCrossAliyunId(v string) *CreateRestoreTaskRequest {
	s.CrossAliyunId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetCrossRoleName(v string) *CreateRestoreTaskRequest {
	s.CrossRoleName = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetOwnerId(v string) *CreateRestoreTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetClientToken(v string) *CreateRestoreTaskRequest {
	s.ClientToken = &v
	return s
}

type CreateRestoreTaskResponseBody struct {
	RestoreTaskId  *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s CreateRestoreTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRestoreTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestoreTaskResponseBody) SetRestoreTaskId(v string) *CreateRestoreTaskResponseBody {
	s.RestoreTaskId = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetRequestId(v string) *CreateRestoreTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetHttpStatusCode(v int32) *CreateRestoreTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetErrMessage(v string) *CreateRestoreTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetSuccess(v bool) *CreateRestoreTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetErrCode(v string) *CreateRestoreTaskResponseBody {
	s.ErrCode = &v
	return s
}

type CreateRestoreTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRestoreTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRestoreTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRestoreTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRestoreTaskResponse) SetHeaders(v map[string]*string) *CreateRestoreTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRestoreTaskResponse) SetBody(v *CreateRestoreTaskResponseBody) *CreateRestoreTaskResponse {
	s.Body = v
	return s
}

type DescribeBackupGatewayListRequest struct {
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Identifier  *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeBackupGatewayListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupGatewayListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListRequest) SetRegion(v string) *DescribeBackupGatewayListRequest {
	s.Region = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetIdentifier(v string) *DescribeBackupGatewayListRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetPageSize(v int32) *DescribeBackupGatewayListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetPageNum(v int32) *DescribeBackupGatewayListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetOwnerId(v string) *DescribeBackupGatewayListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupGatewayListRequest) SetClientToken(v string) *DescribeBackupGatewayListRequest {
	s.ClientToken = &v
	return s
}

type DescribeBackupGatewayListResponseBody struct {
	PageNum        *int32                                      `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	TotalElements  *int32                                      `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	TotalPages     *int32                                      `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Items          *DescribeBackupGatewayListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	ErrMessage     *string                                     `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                     `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeBackupGatewayListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupGatewayListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListResponseBody) SetPageNum(v int32) *DescribeBackupGatewayListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetRequestId(v string) *DescribeBackupGatewayListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetPageSize(v int32) *DescribeBackupGatewayListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetHttpStatusCode(v int32) *DescribeBackupGatewayListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetTotalElements(v int32) *DescribeBackupGatewayListResponseBody {
	s.TotalElements = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetTotalPages(v int32) *DescribeBackupGatewayListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetItems(v *DescribeBackupGatewayListResponseBodyItems) *DescribeBackupGatewayListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetErrMessage(v string) *DescribeBackupGatewayListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetSuccess(v bool) *DescribeBackupGatewayListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBody) SetErrCode(v string) *DescribeBackupGatewayListResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeBackupGatewayListResponseBodyItems struct {
	BackupGateway []*DescribeBackupGatewayListResponseBodyItemsBackupGateway `json:"BackupGateway,omitempty" xml:"BackupGateway,omitempty" type:"Repeated"`
}

func (s DescribeBackupGatewayListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupGatewayListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListResponseBodyItems) SetBackupGateway(v []*DescribeBackupGatewayListResponseBodyItemsBackupGateway) *DescribeBackupGatewayListResponseBodyItems {
	s.BackupGateway = v
	return s
}

type DescribeBackupGatewayListResponseBodyItemsBackupGateway struct {
	DisplayName              *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	BackupGatewayCreateTime  *int64  `json:"BackupGatewayCreateTime,omitempty" xml:"BackupGatewayCreateTime,omitempty"`
	BackupGatewayId          *string `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	Region                   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Identifier               *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	SourceEndpointInternetIP *string `json:"SourceEndpointInternetIP,omitempty" xml:"SourceEndpointInternetIP,omitempty"`
	BackupGatewayStatus      *string `json:"BackupGatewayStatus,omitempty" xml:"BackupGatewayStatus,omitempty"`
	SourceEndpointIntranetIP *string `json:"SourceEndpointIntranetIP,omitempty" xml:"SourceEndpointIntranetIP,omitempty"`
	LastHeartbeatTime        *int64  `json:"LastHeartbeatTime,omitempty" xml:"LastHeartbeatTime,omitempty"`
	SourceEndpointHostname   *string `json:"SourceEndpointHostname,omitempty" xml:"SourceEndpointHostname,omitempty"`
}

func (s DescribeBackupGatewayListResponseBodyItemsBackupGateway) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupGatewayListResponseBodyItemsBackupGateway) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetDisplayName(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.DisplayName = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetBackupGatewayCreateTime(v int64) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.BackupGatewayCreateTime = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetBackupGatewayId(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.BackupGatewayId = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetRegion(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.Region = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetIdentifier(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.Identifier = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetSourceEndpointInternetIP(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.SourceEndpointInternetIP = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetBackupGatewayStatus(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.BackupGatewayStatus = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetSourceEndpointIntranetIP(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.SourceEndpointIntranetIP = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetLastHeartbeatTime(v int64) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.LastHeartbeatTime = &v
	return s
}

func (s *DescribeBackupGatewayListResponseBodyItemsBackupGateway) SetSourceEndpointHostname(v string) *DescribeBackupGatewayListResponseBodyItemsBackupGateway {
	s.SourceEndpointHostname = &v
	return s
}

type DescribeBackupGatewayListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupGatewayListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupGatewayListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupGatewayListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListResponse) SetHeaders(v map[string]*string) *DescribeBackupGatewayListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupGatewayListResponse) SetBody(v *DescribeBackupGatewayListResponseBody) *DescribeBackupGatewayListResponse {
	s.Body = v
	return s
}

type DescribeBackupPlanBillingRequest struct {
	BackupPlanId    *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ShowStorageType *bool   `json:"ShowStorageType,omitempty" xml:"ShowStorageType,omitempty"`
}

func (s DescribeBackupPlanBillingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanBillingRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanBillingRequest) SetBackupPlanId(v string) *DescribeBackupPlanBillingRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupPlanBillingRequest) SetOwnerId(v string) *DescribeBackupPlanBillingRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPlanBillingRequest) SetClientToken(v string) *DescribeBackupPlanBillingRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeBackupPlanBillingRequest) SetShowStorageType(v bool) *DescribeBackupPlanBillingRequest {
	s.ShowStorageType = &v
	return s
}

type DescribeBackupPlanBillingResponseBody struct {
	Item           *DescribeBackupPlanBillingResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Struct"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage     *string                                    `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                    `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeBackupPlanBillingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanBillingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanBillingResponseBody) SetItem(v *DescribeBackupPlanBillingResponseBodyItem) *DescribeBackupPlanBillingResponseBody {
	s.Item = v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetRequestId(v string) *DescribeBackupPlanBillingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetHttpStatusCode(v int32) *DescribeBackupPlanBillingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetErrMessage(v string) *DescribeBackupPlanBillingResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetSuccess(v bool) *DescribeBackupPlanBillingResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBody) SetErrCode(v string) *DescribeBackupPlanBillingResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeBackupPlanBillingResponseBodyItem struct {
	FullStorageSize      *int64  `json:"FullStorageSize,omitempty" xml:"FullStorageSize,omitempty"`
	BuyCreateTimestamp   *int64  `json:"BuyCreateTimestamp,omitempty" xml:"BuyCreateTimestamp,omitempty"`
	BuyExpiredTimestamp  *int64  `json:"BuyExpiredTimestamp,omitempty" xml:"BuyExpiredTimestamp,omitempty"`
	QuotaEndTimestamp    *int64  `json:"QuotaEndTimestamp,omitempty" xml:"QuotaEndTimestamp,omitempty"`
	BuySpec              *string `json:"BuySpec,omitempty" xml:"BuySpec,omitempty"`
	QuotaStartTimestamp  *int64  `json:"QuotaStartTimestamp,omitempty" xml:"QuotaStartTimestamp,omitempty"`
	IsExpired            *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	PaiedBytes           *int64  `json:"PaiedBytes,omitempty" xml:"PaiedBytes,omitempty"`
	TotalFreeBytes       *int64  `json:"TotalFreeBytes,omitempty" xml:"TotalFreeBytes,omitempty"`
	UsedFullBytes        *int64  `json:"UsedFullBytes,omitempty" xml:"UsedFullBytes,omitempty"`
	IsFreeBytesUnlimited *bool   `json:"IsFreeBytesUnlimited,omitempty" xml:"IsFreeBytesUnlimited,omitempty"`
	ContStorageSize      *int64  `json:"ContStorageSize,omitempty" xml:"ContStorageSize,omitempty"`
	BuyChargeType        *string `json:"BuyChargeType,omitempty" xml:"BuyChargeType,omitempty"`
	UsedIncrementBytes   *int64  `json:"UsedIncrementBytes,omitempty" xml:"UsedIncrementBytes,omitempty"`
}

func (s DescribeBackupPlanBillingResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanBillingResponseBodyItem) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetFullStorageSize(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.FullStorageSize = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetBuyCreateTimestamp(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.BuyCreateTimestamp = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetBuyExpiredTimestamp(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.BuyExpiredTimestamp = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetQuotaEndTimestamp(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.QuotaEndTimestamp = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetBuySpec(v string) *DescribeBackupPlanBillingResponseBodyItem {
	s.BuySpec = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetQuotaStartTimestamp(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.QuotaStartTimestamp = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetIsExpired(v bool) *DescribeBackupPlanBillingResponseBodyItem {
	s.IsExpired = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetPaiedBytes(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.PaiedBytes = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetTotalFreeBytes(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.TotalFreeBytes = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetUsedFullBytes(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.UsedFullBytes = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetIsFreeBytesUnlimited(v bool) *DescribeBackupPlanBillingResponseBodyItem {
	s.IsFreeBytesUnlimited = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetContStorageSize(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.ContStorageSize = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetBuyChargeType(v string) *DescribeBackupPlanBillingResponseBodyItem {
	s.BuyChargeType = &v
	return s
}

func (s *DescribeBackupPlanBillingResponseBodyItem) SetUsedIncrementBytes(v int64) *DescribeBackupPlanBillingResponseBodyItem {
	s.UsedIncrementBytes = &v
	return s
}

type DescribeBackupPlanBillingResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPlanBillingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPlanBillingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanBillingResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanBillingResponse) SetHeaders(v map[string]*string) *DescribeBackupPlanBillingResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlanBillingResponse) SetBody(v *DescribeBackupPlanBillingResponseBody) *DescribeBackupPlanBillingResponse {
	s.Body = v
	return s
}

type DescribeBackupPlanListRequest struct {
	BackupPlanId     *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	BackupPlanName   *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	BackupPlanStatus *string `json:"BackupPlanStatus,omitempty" xml:"BackupPlanStatus,omitempty"`
	OwnerId          *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeBackupPlanListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListRequest) SetBackupPlanId(v string) *DescribeBackupPlanListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetRegion(v string) *DescribeBackupPlanListRequest {
	s.Region = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetPageSize(v int32) *DescribeBackupPlanListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetPageNum(v int32) *DescribeBackupPlanListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetBackupPlanName(v string) *DescribeBackupPlanListRequest {
	s.BackupPlanName = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetBackupPlanStatus(v string) *DescribeBackupPlanListRequest {
	s.BackupPlanStatus = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetOwnerId(v string) *DescribeBackupPlanListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPlanListRequest) SetClientToken(v string) *DescribeBackupPlanListRequest {
	s.ClientToken = &v
	return s
}

type DescribeBackupPlanListResponseBody struct {
	PageNum        *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	TotalElements  *int32                                   `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	TotalPages     *int32                                   `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Items          *DescribeBackupPlanListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	ErrMessage     *string                                  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeBackupPlanListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListResponseBody) SetPageNum(v int32) *DescribeBackupPlanListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetRequestId(v string) *DescribeBackupPlanListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetPageSize(v int32) *DescribeBackupPlanListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetHttpStatusCode(v int32) *DescribeBackupPlanListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetTotalElements(v int32) *DescribeBackupPlanListResponseBody {
	s.TotalElements = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetTotalPages(v int32) *DescribeBackupPlanListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetItems(v *DescribeBackupPlanListResponseBodyItems) *DescribeBackupPlanListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetErrMessage(v string) *DescribeBackupPlanListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetSuccess(v bool) *DescribeBackupPlanListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetErrCode(v string) *DescribeBackupPlanListResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeBackupPlanListResponseBodyItems struct {
	BackupPlanDetail []*DescribeBackupPlanListResponseBodyItemsBackupPlanDetail `json:"BackupPlanDetail,omitempty" xml:"BackupPlanDetail,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlanListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListResponseBodyItems) SetBackupPlanDetail(v []*DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) *DescribeBackupPlanListResponseBodyItems {
	s.BackupPlanDetail = v
	return s
}

type DescribeBackupPlanListResponseBodyItemsBackupPlanDetail struct {
	CrossRoleName                        *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	SourceEndpointInstanceType           *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	BackupSetDownloadDir                 *string `json:"BackupSetDownloadDir,omitempty" xml:"BackupSetDownloadDir,omitempty"`
	SourceEndpointIpPort                 *string `json:"SourceEndpointIpPort,omitempty" xml:"SourceEndpointIpPort,omitempty"`
	CrossAliyunId                        *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	DuplicationArchivePeriod             *int32  `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	BackupPlanId                         *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	EndTimestampForRestore               *int64  `json:"EndTimestampForRestore,omitempty" xml:"EndTimestampForRestore,omitempty"`
	BackupPlanStatus                     *string `json:"BackupPlanStatus,omitempty" xml:"BackupPlanStatus,omitempty"`
	BackupSetDownloadFullDataFormat      *string `json:"BackupSetDownloadFullDataFormat,omitempty" xml:"BackupSetDownloadFullDataFormat,omitempty"`
	BackupRetentionPeriod                *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	OSSBucketRegion                      *string `json:"OSSBucketRegion,omitempty" xml:"OSSBucketRegion,omitempty"`
	SourceEndpointOracleSID              *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	BackupStorageType                    *string `json:"BackupStorageType,omitempty" xml:"BackupStorageType,omitempty"`
	BackupMethod                         *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	SourceEndpointRegion                 *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	SourceEndpointDatabaseName           *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	BackupPeriod                         *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupSetDownloadGatewayId           *int64  `json:"BackupSetDownloadGatewayId,omitempty" xml:"BackupSetDownloadGatewayId,omitempty"`
	InstanceClass                        *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	BackupPlanCreateTime                 *int64  `json:"BackupPlanCreateTime,omitempty" xml:"BackupPlanCreateTime,omitempty"`
	BackupStartTime                      *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	DuplicationInfrequentAccessPeriod    *int32  `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	BackupSetDownloadTargetType          *string `json:"BackupSetDownloadTargetType,omitempty" xml:"BackupSetDownloadTargetType,omitempty"`
	ErrMessage                           *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	BackupObjects                        *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	BeginTimestampForRestore             *int64  `json:"BeginTimestampForRestore,omitempty" xml:"BeginTimestampForRestore,omitempty"`
	SourceEndpointInstanceID             *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	OpenBackupSetAutoDownload            *bool   `json:"OpenBackupSetAutoDownload,omitempty" xml:"OpenBackupSetAutoDownload,omitempty"`
	BackupPlanName                       *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	OSSBucketName                        *string `json:"OSSBucketName,omitempty" xml:"OSSBucketName,omitempty"`
	BackupGatewayId                      *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	SourceEndpointUserName               *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	BackupSetDownloadIncrementDataFormat *string `json:"BackupSetDownloadIncrementDataFormat,omitempty" xml:"BackupSetDownloadIncrementDataFormat,omitempty"`
	EnableBackupLog                      *bool   `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
}

func (s DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetCrossRoleName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.CrossRoleName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointInstanceType(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadDir(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadDir = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointIpPort(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointIpPort = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetCrossAliyunId(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.CrossAliyunId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetDuplicationArchivePeriod(v int32) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.DuplicationArchivePeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanId(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetEndTimestampForRestore(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.EndTimestampForRestore = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanStatus(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanStatus = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadFullDataFormat(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadFullDataFormat = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupRetentionPeriod(v int32) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetOSSBucketRegion(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.OSSBucketRegion = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointOracleSID(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupStorageType(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupStorageType = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupMethod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointRegion(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointDatabaseName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPeriod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadGatewayId(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadGatewayId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetInstanceClass(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.InstanceClass = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanCreateTime(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanCreateTime = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupStartTime(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetDuplicationInfrequentAccessPeriod(v int32) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.DuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadTargetType(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadTargetType = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetErrMessage(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupObjects(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupObjects = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBeginTimestampForRestore(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BeginTimestampForRestore = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointInstanceID(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetOpenBackupSetAutoDownload(v bool) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.OpenBackupSetAutoDownload = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetOSSBucketName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.OSSBucketName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupGatewayId(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupGatewayId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointUserName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointUserName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadIncrementDataFormat(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadIncrementDataFormat = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetEnableBackupLog(v bool) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.EnableBackupLog = &v
	return s
}

type DescribeBackupPlanListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPlanListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPlanListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPlanListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListResponse) SetHeaders(v map[string]*string) *DescribeBackupPlanListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPlanListResponse) SetBody(v *DescribeBackupPlanListResponseBody) *DescribeBackupPlanListResponse {
	s.Body = v
	return s
}

type DescribeBackupSetDownloadTaskListRequest struct {
	BackupPlanId            *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupSetDownloadTaskId *string `json:"BackupSetDownloadTaskId,omitempty" xml:"BackupSetDownloadTaskId,omitempty"`
	PageSize                *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum                 *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	OwnerId                 *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeBackupSetDownloadTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetBackupPlanId(v string) *DescribeBackupSetDownloadTaskListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetBackupSetDownloadTaskId(v string) *DescribeBackupSetDownloadTaskListRequest {
	s.BackupSetDownloadTaskId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetPageSize(v int32) *DescribeBackupSetDownloadTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetPageNum(v int32) *DescribeBackupSetDownloadTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetOwnerId(v string) *DescribeBackupSetDownloadTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListRequest) SetClientToken(v string) *DescribeBackupSetDownloadTaskListRequest {
	s.ClientToken = &v
	return s
}

type DescribeBackupSetDownloadTaskListResponseBody struct {
	PageNum        *int32                                              `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HttpStatusCode *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	TotalElements  *int32                                              `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	TotalPages     *int32                                              `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Items          *DescribeBackupSetDownloadTaskListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	ErrMessage     *string                                             `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeBackupSetDownloadTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetPageNum(v int32) *DescribeBackupSetDownloadTaskListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetRequestId(v string) *DescribeBackupSetDownloadTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetPageSize(v int32) *DescribeBackupSetDownloadTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetHttpStatusCode(v int32) *DescribeBackupSetDownloadTaskListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetTotalElements(v int32) *DescribeBackupSetDownloadTaskListResponseBody {
	s.TotalElements = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetTotalPages(v int32) *DescribeBackupSetDownloadTaskListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetItems(v *DescribeBackupSetDownloadTaskListResponseBodyItems) *DescribeBackupSetDownloadTaskListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetErrMessage(v string) *DescribeBackupSetDownloadTaskListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetSuccess(v bool) *DescribeBackupSetDownloadTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetErrCode(v string) *DescribeBackupSetDownloadTaskListResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeBackupSetDownloadTaskListResponseBodyItems struct {
	BackupSetDownloadTaskDetail []*DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail `json:"BackupSetDownloadTaskDetail,omitempty" xml:"BackupSetDownloadTaskDetail,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetDownloadTaskListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItems) SetBackupSetDownloadTaskDetail(v []*DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) *DescribeBackupSetDownloadTaskListResponseBodyItems {
	s.BackupSetDownloadTaskDetail = v
	return s
}

type DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail struct {
	BackupSetDownloadStatus      *string `json:"BackupSetDownloadStatus,omitempty" xml:"BackupSetDownloadStatus,omitempty"`
	BackupSetDataSize            *int64  `json:"BackupSetDataSize,omitempty" xml:"BackupSetDataSize,omitempty"`
	BackupSetDownloadDir         *string `json:"BackupSetDownloadDir,omitempty" xml:"BackupSetDownloadDir,omitempty"`
	BackupSetDownloadTargetType  *string `json:"BackupSetDownloadTargetType,omitempty" xml:"BackupSetDownloadTargetType,omitempty"`
	ErrMessage                   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	BackupPlanId                 *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupSetDownloadCreateTime  *int64  `json:"BackupSetDownloadCreateTime,omitempty" xml:"BackupSetDownloadCreateTime,omitempty"`
	BackupSetDbType              *string `json:"BackupSetDbType,omitempty" xml:"BackupSetDbType,omitempty"`
	BackupSetDownloadInternetUrl *string `json:"BackupSetDownloadInternetUrl,omitempty" xml:"BackupSetDownloadInternetUrl,omitempty"`
	BackupSetId                  *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupGatewayId              *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	BackupSetDownloadIntranetUrl *string `json:"BackupSetDownloadIntranetUrl,omitempty" xml:"BackupSetDownloadIntranetUrl,omitempty"`
	BackupSetDownloadWay         *string `json:"BackupSetDownloadWay,omitempty" xml:"BackupSetDownloadWay,omitempty"`
	BackupSetDownloadFinishTime  *int64  `json:"BackupSetDownloadFinishTime,omitempty" xml:"BackupSetDownloadFinishTime,omitempty"`
	BackupSetJobType             *string `json:"BackupSetJobType,omitempty" xml:"BackupSetJobType,omitempty"`
	BackupSetDownloadTaskId      *string `json:"BackupSetDownloadTaskId,omitempty" xml:"BackupSetDownloadTaskId,omitempty"`
	BackupSetDownloadTaskName    *string `json:"BackupSetDownloadTaskName,omitempty" xml:"BackupSetDownloadTaskName,omitempty"`
	BackupSetDataFormat          *string `json:"BackupSetDataFormat,omitempty" xml:"BackupSetDataFormat,omitempty"`
	BackupSetCode                *string `json:"BackupSetCode,omitempty" xml:"BackupSetCode,omitempty"`
}

func (s DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadStatus(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadStatus = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDataSize(v int64) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDataSize = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadDir(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadDir = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadTargetType(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadTargetType = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetErrMessage(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupPlanId(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadCreateTime(v int64) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadCreateTime = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDbType(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDbType = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadInternetUrl(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadInternetUrl = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetId(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupGatewayId(v int64) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupGatewayId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadIntranetUrl(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadIntranetUrl = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadWay(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadWay = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadFinishTime(v int64) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadFinishTime = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetJobType(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetJobType = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadTaskId(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadTaskId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadTaskName(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadTaskName = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDataFormat(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDataFormat = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetCode(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetCode = &v
	return s
}

type DescribeBackupSetDownloadTaskListResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupSetDownloadTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupSetDownloadTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListResponse) SetHeaders(v map[string]*string) *DescribeBackupSetDownloadTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponse) SetBody(v *DescribeBackupSetDownloadTaskListResponseBody) *DescribeBackupSetDownloadTaskListResponse {
	s.Body = v
	return s
}

type DescribeFullBackupListRequest struct {
	BackupPlanId    *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum         *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ShowStorageType *bool   `json:"ShowStorageType,omitempty" xml:"ShowStorageType,omitempty"`
	BackupSetId     *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
}

func (s DescribeFullBackupListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullBackupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListRequest) SetBackupPlanId(v string) *DescribeFullBackupListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetPageSize(v int32) *DescribeFullBackupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetPageNum(v int32) *DescribeFullBackupListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetOwnerId(v string) *DescribeFullBackupListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetClientToken(v string) *DescribeFullBackupListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetShowStorageType(v bool) *DescribeFullBackupListRequest {
	s.ShowStorageType = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetBackupSetId(v string) *DescribeFullBackupListRequest {
	s.BackupSetId = &v
	return s
}

type DescribeFullBackupListResponseBody struct {
	PageNum        *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	TotalElements  *int32                                   `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	TotalPages     *int32                                   `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Items          *DescribeFullBackupListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	ErrMessage     *string                                  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeFullBackupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullBackupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListResponseBody) SetPageNum(v int32) *DescribeFullBackupListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetRequestId(v string) *DescribeFullBackupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetPageSize(v int32) *DescribeFullBackupListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetHttpStatusCode(v int32) *DescribeFullBackupListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetTotalElements(v int32) *DescribeFullBackupListResponseBody {
	s.TotalElements = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetTotalPages(v int32) *DescribeFullBackupListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetItems(v *DescribeFullBackupListResponseBodyItems) *DescribeFullBackupListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetErrMessage(v string) *DescribeFullBackupListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetSuccess(v bool) *DescribeFullBackupListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetErrCode(v string) *DescribeFullBackupListResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeFullBackupListResponseBodyItems struct {
	FullBackupFile []*DescribeFullBackupListResponseBodyItemsFullBackupFile `json:"FullBackupFile,omitempty" xml:"FullBackupFile,omitempty" type:"Repeated"`
}

func (s DescribeFullBackupListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullBackupListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListResponseBodyItems) SetFullBackupFile(v []*DescribeFullBackupListResponseBodyItemsFullBackupFile) *DescribeFullBackupListResponseBodyItems {
	s.FullBackupFile = v
	return s
}

type DescribeFullBackupListResponseBodyItemsFullBackupFile struct {
	FinishTime           *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	BackupStatus         *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SourceEndpointIpPort *string `json:"SourceEndpointIpPort,omitempty" xml:"SourceEndpointIpPort,omitempty"`
	BackupObjects        *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	ErrMessage           *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BackupSetExpiredTime *int64  `json:"BackupSetExpiredTime,omitempty" xml:"BackupSetExpiredTime,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StorageMethod        *string `json:"StorageMethod,omitempty" xml:"StorageMethod,omitempty"`
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSize           *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
}

func (s DescribeFullBackupListResponseBodyItemsFullBackupFile) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullBackupListResponseBodyItemsFullBackupFile) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetFinishTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.FinishTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupStatus(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupStatus = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetCreateTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.CreateTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointIpPort(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointIpPort = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupObjects(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupObjects = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetErrMessage(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.ErrMessage = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetEndTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.EndTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupSetExpiredTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupSetExpiredTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetStartTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.StartTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetStorageMethod(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.StorageMethod = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupSetId(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupSetId = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupSize(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupSize = &v
	return s
}

type DescribeFullBackupListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFullBackupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFullBackupListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullBackupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListResponse) SetHeaders(v map[string]*string) *DescribeFullBackupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFullBackupListResponse) SetBody(v *DescribeFullBackupListResponseBody) *DescribeFullBackupListResponse {
	s.Body = v
	return s
}

type DescribeIncrementBackupListRequest struct {
	BackupPlanId    *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum         *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ShowStorageType *bool   `json:"ShowStorageType,omitempty" xml:"ShowStorageType,omitempty"`
}

func (s DescribeIncrementBackupListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIncrementBackupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListRequest) SetBackupPlanId(v string) *DescribeIncrementBackupListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetPageSize(v int32) *DescribeIncrementBackupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetPageNum(v int32) *DescribeIncrementBackupListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetOwnerId(v string) *DescribeIncrementBackupListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetClientToken(v string) *DescribeIncrementBackupListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeIncrementBackupListRequest) SetShowStorageType(v bool) *DescribeIncrementBackupListRequest {
	s.ShowStorageType = &v
	return s
}

type DescribeIncrementBackupListResponseBody struct {
	PageNum        *int32                                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	TotalElements  *int32                                        `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	TotalPages     *int32                                        `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Items          *DescribeIncrementBackupListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	ErrMessage     *string                                       `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                       `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeIncrementBackupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIncrementBackupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListResponseBody) SetPageNum(v int32) *DescribeIncrementBackupListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetRequestId(v string) *DescribeIncrementBackupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetPageSize(v int32) *DescribeIncrementBackupListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetHttpStatusCode(v int32) *DescribeIncrementBackupListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetTotalElements(v int32) *DescribeIncrementBackupListResponseBody {
	s.TotalElements = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetTotalPages(v int32) *DescribeIncrementBackupListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetItems(v *DescribeIncrementBackupListResponseBodyItems) *DescribeIncrementBackupListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetErrMessage(v string) *DescribeIncrementBackupListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetSuccess(v bool) *DescribeIncrementBackupListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetErrCode(v string) *DescribeIncrementBackupListResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeIncrementBackupListResponseBodyItems struct {
	IncrementBackupFile []*DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile `json:"IncrementBackupFile,omitempty" xml:"IncrementBackupFile,omitempty" type:"Repeated"`
}

func (s DescribeIncrementBackupListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeIncrementBackupListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListResponseBodyItems) SetIncrementBackupFile(v []*DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) *DescribeIncrementBackupListResponseBodyItems {
	s.IncrementBackupFile = v
	return s
}

type DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile struct {
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BackupSetExpiredTime *int64  `json:"BackupSetExpiredTime,omitempty" xml:"BackupSetExpiredTime,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StorageMethod        *string `json:"StorageMethod,omitempty" xml:"StorageMethod,omitempty"`
	BackupSetJobId       *string `json:"BackupSetJobId,omitempty" xml:"BackupSetJobId,omitempty"`
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupStatus         *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	SourceEndpointIpPort *string `json:"SourceEndpointIpPort,omitempty" xml:"SourceEndpointIpPort,omitempty"`
	BackupSize           *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
}

func (s DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) String() string {
	return tea.Prettify(s)
}

func (s DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetEndTime(v int64) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.EndTime = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupSetExpiredTime(v int64) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupSetExpiredTime = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetStartTime(v int64) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.StartTime = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetStorageMethod(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.StorageMethod = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupSetJobId(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupSetJobId = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupSetId(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupSetId = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupStatus(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupStatus = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetSourceEndpointIpPort(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.SourceEndpointIpPort = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupSize(v int64) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupSize = &v
	return s
}

type DescribeIncrementBackupListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIncrementBackupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIncrementBackupListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIncrementBackupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListResponse) SetHeaders(v map[string]*string) *DescribeIncrementBackupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeIncrementBackupListResponse) SetBody(v *DescribeIncrementBackupListResponseBody) *DescribeIncrementBackupListResponse {
	s.Body = v
	return s
}

type DescribeJobErrorCodeRequest struct {
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeJobErrorCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobErrorCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobErrorCodeRequest) SetTaskId(v string) *DescribeJobErrorCodeRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeJobErrorCodeRequest) SetLanguage(v string) *DescribeJobErrorCodeRequest {
	s.Language = &v
	return s
}

func (s *DescribeJobErrorCodeRequest) SetOwnerId(v string) *DescribeJobErrorCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeJobErrorCodeRequest) SetClientToken(v string) *DescribeJobErrorCodeRequest {
	s.ClientToken = &v
	return s
}

type DescribeJobErrorCodeResponseBody struct {
	Item           *DescribeJobErrorCodeResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Struct"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage     *string                               `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                               `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeJobErrorCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobErrorCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobErrorCodeResponseBody) SetItem(v *DescribeJobErrorCodeResponseBodyItem) *DescribeJobErrorCodeResponseBody {
	s.Item = v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetRequestId(v string) *DescribeJobErrorCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetHttpStatusCode(v int32) *DescribeJobErrorCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetErrMessage(v string) *DescribeJobErrorCodeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetSuccess(v bool) *DescribeJobErrorCodeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetErrCode(v string) *DescribeJobErrorCodeResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeJobErrorCodeResponseBodyItem struct {
	JobState     *string `json:"JobState,omitempty" xml:"JobState,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType      *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Language     *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s DescribeJobErrorCodeResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobErrorCodeResponseBodyItem) GoString() string {
	return s.String()
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetJobState(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.JobState = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetErrorMessage(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetErrorCode(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.ErrorCode = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetJobId(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.JobId = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetJobType(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.JobType = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetLanguage(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.Language = &v
	return s
}

type DescribeJobErrorCodeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeJobErrorCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobErrorCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobErrorCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobErrorCodeResponse) SetHeaders(v map[string]*string) *DescribeJobErrorCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobErrorCodeResponse) SetBody(v *DescribeJobErrorCodeResponseBody) *DescribeJobErrorCodeResponse {
	s.Body = v
	return s
}

type DescribeNodeCidrListRequest struct {
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeNodeCidrListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeCidrListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListRequest) SetRegion(v string) *DescribeNodeCidrListRequest {
	s.Region = &v
	return s
}

func (s *DescribeNodeCidrListRequest) SetOwnerId(v string) *DescribeNodeCidrListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNodeCidrListRequest) SetClientToken(v string) *DescribeNodeCidrListRequest {
	s.ClientToken = &v
	return s
}

type DescribeNodeCidrListResponseBody struct {
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InternetIPs    *DescribeNodeCidrListResponseBodyInternetIPs `json:"InternetIPs,omitempty" xml:"InternetIPs,omitempty" type:"Struct"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	IntranetIPs    *DescribeNodeCidrListResponseBodyIntranetIPs `json:"IntranetIPs,omitempty" xml:"IntranetIPs,omitempty" type:"Struct"`
	ErrMessage     *string                                      `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                      `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeNodeCidrListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeCidrListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListResponseBody) SetRequestId(v string) *DescribeNodeCidrListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetInternetIPs(v *DescribeNodeCidrListResponseBodyInternetIPs) *DescribeNodeCidrListResponseBody {
	s.InternetIPs = v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetHttpStatusCode(v int32) *DescribeNodeCidrListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetIntranetIPs(v *DescribeNodeCidrListResponseBodyIntranetIPs) *DescribeNodeCidrListResponseBody {
	s.IntranetIPs = v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetErrMessage(v string) *DescribeNodeCidrListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetSuccess(v bool) *DescribeNodeCidrListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNodeCidrListResponseBody) SetErrCode(v string) *DescribeNodeCidrListResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeNodeCidrListResponseBodyInternetIPs struct {
	InternetIP []*string `json:"InternetIP,omitempty" xml:"InternetIP,omitempty" type:"Repeated"`
}

func (s DescribeNodeCidrListResponseBodyInternetIPs) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeCidrListResponseBodyInternetIPs) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListResponseBodyInternetIPs) SetInternetIP(v []*string) *DescribeNodeCidrListResponseBodyInternetIPs {
	s.InternetIP = v
	return s
}

type DescribeNodeCidrListResponseBodyIntranetIPs struct {
	IntranetIP []*string `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty" type:"Repeated"`
}

func (s DescribeNodeCidrListResponseBodyIntranetIPs) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeCidrListResponseBodyIntranetIPs) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListResponseBodyIntranetIPs) SetIntranetIP(v []*string) *DescribeNodeCidrListResponseBodyIntranetIPs {
	s.IntranetIP = v
	return s
}

type DescribeNodeCidrListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNodeCidrListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeCidrListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeCidrListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListResponse) SetHeaders(v map[string]*string) *DescribeNodeCidrListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeCidrListResponse) SetBody(v *DescribeNodeCidrListResponseBody) *DescribeNodeCidrListResponse {
	s.Body = v
	return s
}

type DescribePreCheckProgressListRequest struct {
	BackupPlanId  *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribePreCheckProgressListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckProgressListRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListRequest) SetBackupPlanId(v string) *DescribePreCheckProgressListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribePreCheckProgressListRequest) SetRestoreTaskId(v string) *DescribePreCheckProgressListRequest {
	s.RestoreTaskId = &v
	return s
}

func (s *DescribePreCheckProgressListRequest) SetOwnerId(v string) *DescribePreCheckProgressListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePreCheckProgressListRequest) SetClientToken(v string) *DescribePreCheckProgressListRequest {
	s.ClientToken = &v
	return s
}

type DescribePreCheckProgressListResponseBody struct {
	Status         *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Progress       *int32                                         `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          *DescribePreCheckProgressListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	ErrMessage     *string                                        `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                        `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribePreCheckProgressListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckProgressListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListResponseBody) SetStatus(v string) *DescribePreCheckProgressListResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetProgress(v int32) *DescribePreCheckProgressListResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetRequestId(v string) *DescribePreCheckProgressListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetHttpStatusCode(v int32) *DescribePreCheckProgressListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetItems(v *DescribePreCheckProgressListResponseBodyItems) *DescribePreCheckProgressListResponseBody {
	s.Items = v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetErrMessage(v string) *DescribePreCheckProgressListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetSuccess(v bool) *DescribePreCheckProgressListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBody) SetErrCode(v string) *DescribePreCheckProgressListResponseBody {
	s.ErrCode = &v
	return s
}

type DescribePreCheckProgressListResponseBodyItems struct {
	PreCheckProgressDetail []*DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail `json:"PreCheckProgressDetail,omitempty" xml:"PreCheckProgressDetail,omitempty" type:"Repeated"`
}

func (s DescribePreCheckProgressListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckProgressListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListResponseBodyItems) SetPreCheckProgressDetail(v []*DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) *DescribePreCheckProgressListResponseBodyItems {
	s.PreCheckProgressDetail = v
	return s
}

type DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail struct {
	FinishTime *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	BootTime   *int64  `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Item       *string `json:"Item,omitempty" xml:"Item,omitempty"`
	ErrMsg     *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	OrderNum   *string `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	Names      *string `json:"Names,omitempty" xml:"Names,omitempty"`
}

func (s DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetFinishTime(v int64) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.FinishTime = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetState(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.State = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetBootTime(v int64) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.BootTime = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetJobId(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetItem(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.Item = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetErrMsg(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetOrderNum(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.OrderNum = &v
	return s
}

func (s *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail) SetNames(v string) *DescribePreCheckProgressListResponseBodyItemsPreCheckProgressDetail {
	s.Names = &v
	return s
}

type DescribePreCheckProgressListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePreCheckProgressListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePreCheckProgressListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckProgressListResponse) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListResponse) SetHeaders(v map[string]*string) *DescribePreCheckProgressListResponse {
	s.Headers = v
	return s
}

func (s *DescribePreCheckProgressListResponse) SetBody(v *DescribePreCheckProgressListResponseBody) *DescribePreCheckProgressListResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerId(v string) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetClientToken(v string) *DescribeRegionsRequest {
	s.ClientToken = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Regions        *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	ErrMessage     *string                             `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetHttpStatusCode(v int32) *DescribeRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrMessage(v string) *DescribeRegionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrCode(v string) *DescribeRegionsResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionCode []*string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionCode(v []*string) *DescribeRegionsResponseBodyRegions {
	s.RegionCode = v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRestoreRangeInfoRequest struct {
	BackupPlanId             *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BeginTimestampForRestore *int64  `json:"BeginTimestampForRestore,omitempty" xml:"BeginTimestampForRestore,omitempty"`
	EndTimestampForRestore   *int64  `json:"EndTimestampForRestore,omitempty" xml:"EndTimestampForRestore,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RecentlyRestore          *bool   `json:"RecentlyRestore,omitempty" xml:"RecentlyRestore,omitempty"`
}

func (s DescribeRestoreRangeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreRangeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoRequest) SetBackupPlanId(v string) *DescribeRestoreRangeInfoRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetBeginTimestampForRestore(v int64) *DescribeRestoreRangeInfoRequest {
	s.BeginTimestampForRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetEndTimestampForRestore(v int64) *DescribeRestoreRangeInfoRequest {
	s.EndTimestampForRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetOwnerId(v string) *DescribeRestoreRangeInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetClientToken(v string) *DescribeRestoreRangeInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRestoreRangeInfoRequest) SetRecentlyRestore(v bool) *DescribeRestoreRangeInfoRequest {
	s.RecentlyRestore = &v
	return s
}

type DescribeRestoreRangeInfoResponseBody struct {
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          *DescribeRestoreRangeInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	ErrMessage     *string                                    `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                    `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeRestoreRangeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBody) SetRequestId(v string) *DescribeRestoreRangeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetHttpStatusCode(v int32) *DescribeRestoreRangeInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetItems(v *DescribeRestoreRangeInfoResponseBodyItems) *DescribeRestoreRangeInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetErrMessage(v string) *DescribeRestoreRangeInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetSuccess(v bool) *DescribeRestoreRangeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBody) SetErrCode(v string) *DescribeRestoreRangeInfoResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeRestoreRangeInfoResponseBodyItems struct {
	DBSRecoverRange []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange `json:"DBSRecoverRange,omitempty" xml:"DBSRecoverRange,omitempty" type:"Repeated"`
}

func (s DescribeRestoreRangeInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItems) SetDBSRecoverRange(v []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) *DescribeRestoreRangeInfoResponseBodyItems {
	s.DBSRecoverRange = v
	return s
}

type DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange struct {
	EndTimestampForRestore     *int64                                                                  `json:"EndTimestampForRestore,omitempty" xml:"EndTimestampForRestore,omitempty"`
	SourceEndpointInstanceType *string                                                                 `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	FullBackupList             *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList `json:"FullBackupList,omitempty" xml:"FullBackupList,omitempty" type:"Struct"`
	RangeType                  *string                                                                 `json:"RangeType,omitempty" xml:"RangeType,omitempty"`
	BeginTimestampForRestore   *int64                                                                  `json:"BeginTimestampForRestore,omitempty" xml:"BeginTimestampForRestore,omitempty"`
	SourceEndpointInstanceID   *string                                                                 `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetEndTimestampForRestore(v int64) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.EndTimestampForRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetSourceEndpointInstanceType(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetFullBackupList(v *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.FullBackupList = v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetRangeType(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.RangeType = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetBeginTimestampForRestore(v int64) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.BeginTimestampForRestore = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange) SetSourceEndpointInstanceID(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRange {
	s.SourceEndpointInstanceID = &v
	return s
}

type DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList struct {
	FullBackupDetail []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail `json:"FullBackupDetail,omitempty" xml:"FullBackupDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList) SetFullBackupDetail(v []*DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupList {
	s.FullBackupDetail = v
	return s
}

type DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) SetEndTime(v int64) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) SetStartTime(v int64) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail) SetBackupSetId(v string) *DescribeRestoreRangeInfoResponseBodyItemsDBSRecoverRangeFullBackupListFullBackupDetail {
	s.BackupSetId = &v
	return s
}

type DescribeRestoreRangeInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRestoreRangeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreRangeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponse) SetHeaders(v map[string]*string) *DescribeRestoreRangeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreRangeInfoResponse) SetBody(v *DescribeRestoreRangeInfoResponseBody) *DescribeRestoreRangeInfoResponse {
	s.Body = v
	return s
}

type DescribeRestoreTaskListRequest struct {
	BackupPlanId  *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum       *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeRestoreTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListRequest) SetBackupPlanId(v string) *DescribeRestoreTaskListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetRestoreTaskId(v string) *DescribeRestoreTaskListRequest {
	s.RestoreTaskId = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetPageSize(v int32) *DescribeRestoreTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetPageNum(v int32) *DescribeRestoreTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetOwnerId(v string) *DescribeRestoreTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetClientToken(v string) *DescribeRestoreTaskListRequest {
	s.ClientToken = &v
	return s
}

type DescribeRestoreTaskListResponseBody struct {
	PageNum        *int32                                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	TotalElements  *int32                                    `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	TotalPages     *int32                                    `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Items          *DescribeRestoreTaskListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	ErrMessage     *string                                   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string                                   `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeRestoreTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListResponseBody) SetPageNum(v int32) *DescribeRestoreTaskListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetRequestId(v string) *DescribeRestoreTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetPageSize(v int32) *DescribeRestoreTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetHttpStatusCode(v int32) *DescribeRestoreTaskListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetTotalElements(v int32) *DescribeRestoreTaskListResponseBody {
	s.TotalElements = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetTotalPages(v int32) *DescribeRestoreTaskListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetItems(v *DescribeRestoreTaskListResponseBodyItems) *DescribeRestoreTaskListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetErrMessage(v string) *DescribeRestoreTaskListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetSuccess(v bool) *DescribeRestoreTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetErrCode(v string) *DescribeRestoreTaskListResponseBody {
	s.ErrCode = &v
	return s
}

type DescribeRestoreTaskListResponseBodyItems struct {
	RestoreTaskDetail []*DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail `json:"RestoreTaskDetail,omitempty" xml:"RestoreTaskDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTaskListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTaskListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListResponseBodyItems) SetRestoreTaskDetail(v []*DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) *DescribeRestoreTaskListResponseBodyItems {
	s.RestoreTaskDetail = v
	return s
}

type DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail struct {
	RestoreStatus                   *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	FullStruAfterRestoreProgress    *int32  `json:"FullStruAfterRestoreProgress,omitempty" xml:"FullStruAfterRestoreProgress,omitempty"`
	CrossRoleName                   *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	RestoreDir                      *string `json:"RestoreDir,omitempty" xml:"RestoreDir,omitempty"`
	CrossAliyunId                   *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	RestoreObjects                  *string `json:"RestoreObjects,omitempty" xml:"RestoreObjects,omitempty"`
	BackupPlanId                    *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	DestinationEndpointRegion       *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	RestoreTaskCreateTime           *int64  `json:"RestoreTaskCreateTime,omitempty" xml:"RestoreTaskCreateTime,omitempty"`
	DestinationEndpointUserName     *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	RestoreTaskFinishTime           *int64  `json:"RestoreTaskFinishTime,omitempty" xml:"RestoreTaskFinishTime,omitempty"`
	DestinationEndpointIpPort       *string `json:"DestinationEndpointIpPort,omitempty" xml:"DestinationEndpointIpPort,omitempty"`
	DestinationEndpointDatabaseName *string `json:"DestinationEndpointDatabaseName,omitempty" xml:"DestinationEndpointDatabaseName,omitempty"`
	DestinationEndpointOracleSID    *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	FullStruforeRestoreProgress     *int32  `json:"FullStruforeRestoreProgress,omitempty" xml:"FullStruforeRestoreProgress,omitempty"`
	ErrMessage                      *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RestoreTaskId                   *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	FullDataRestoreProgress         *int32  `json:"FullDataRestoreProgress,omitempty" xml:"FullDataRestoreProgress,omitempty"`
	ContinuousRestoreProgress       *int32  `json:"ContinuousRestoreProgress,omitempty" xml:"ContinuousRestoreProgress,omitempty"`
	DestinationEndpointInstanceID   *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	BackupSetId                     *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	RestoreTaskName                 *string `json:"RestoreTaskName,omitempty" xml:"RestoreTaskName,omitempty"`
	BackupGatewayId                 *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	RestoreTime                     *int64  `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreStatus(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreStatus = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetFullStruAfterRestoreProgress(v int32) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.FullStruAfterRestoreProgress = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetCrossRoleName(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.CrossRoleName = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreDir(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreDir = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetCrossAliyunId(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.CrossAliyunId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreObjects(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreObjects = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetBackupPlanId(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointRegion(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTaskCreateTime(v int64) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTaskCreateTime = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointUserName(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTaskFinishTime(v int64) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTaskFinishTime = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointIpPort(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointIpPort = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointDatabaseName(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointDatabaseName = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointOracleSID(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointInstanceType(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetFullStruforeRestoreProgress(v int32) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.FullStruforeRestoreProgress = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetErrMessage(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTaskId(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTaskId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetFullDataRestoreProgress(v int32) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.FullDataRestoreProgress = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetContinuousRestoreProgress(v int32) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.ContinuousRestoreProgress = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointInstanceID(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetBackupSetId(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.BackupSetId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTaskName(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTaskName = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetBackupGatewayId(v int64) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.BackupGatewayId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTime(v int64) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTime = &v
	return s
}

type DescribeRestoreTaskListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRestoreTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRestoreTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRestoreTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListResponse) SetHeaders(v map[string]*string) *DescribeRestoreTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreTaskListResponse) SetBody(v *DescribeRestoreTaskListResponseBody) *DescribeRestoreTaskListResponse {
	s.Body = v
	return s
}

type DisableBackupLogRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DisableBackupLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableBackupLogRequest) GoString() string {
	return s.String()
}

func (s *DisableBackupLogRequest) SetBackupPlanId(v string) *DisableBackupLogRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DisableBackupLogRequest) SetOwnerId(v string) *DisableBackupLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableBackupLogRequest) SetClientToken(v string) *DisableBackupLogRequest {
	s.ClientToken = &v
	return s
}

type DisableBackupLogResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	NeedPrecheck   *bool   `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DisableBackupLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableBackupLogResponseBody) GoString() string {
	return s.String()
}

func (s *DisableBackupLogResponseBody) SetRequestId(v string) *DisableBackupLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetHttpStatusCode(v int32) *DisableBackupLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetBackupPlanId(v string) *DisableBackupLogResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetNeedPrecheck(v bool) *DisableBackupLogResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetErrMessage(v string) *DisableBackupLogResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetSuccess(v bool) *DisableBackupLogResponseBody {
	s.Success = &v
	return s
}

func (s *DisableBackupLogResponseBody) SetErrCode(v string) *DisableBackupLogResponseBody {
	s.ErrCode = &v
	return s
}

type DisableBackupLogResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableBackupLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableBackupLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableBackupLogResponse) GoString() string {
	return s.String()
}

func (s *DisableBackupLogResponse) SetHeaders(v map[string]*string) *DisableBackupLogResponse {
	s.Headers = v
	return s
}

func (s *DisableBackupLogResponse) SetBody(v *DisableBackupLogResponseBody) *DisableBackupLogResponse {
	s.Body = v
	return s
}

type EnableBackupLogRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s EnableBackupLogRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupLogRequest) GoString() string {
	return s.String()
}

func (s *EnableBackupLogRequest) SetBackupPlanId(v string) *EnableBackupLogRequest {
	s.BackupPlanId = &v
	return s
}

func (s *EnableBackupLogRequest) SetOwnerId(v string) *EnableBackupLogRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableBackupLogRequest) SetClientToken(v string) *EnableBackupLogRequest {
	s.ClientToken = &v
	return s
}

type EnableBackupLogResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	NeedPrecheck   *bool   `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s EnableBackupLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupLogResponseBody) GoString() string {
	return s.String()
}

func (s *EnableBackupLogResponseBody) SetRequestId(v string) *EnableBackupLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableBackupLogResponseBody) SetHttpStatusCode(v int32) *EnableBackupLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *EnableBackupLogResponseBody) SetBackupPlanId(v string) *EnableBackupLogResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *EnableBackupLogResponseBody) SetNeedPrecheck(v bool) *EnableBackupLogResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *EnableBackupLogResponseBody) SetErrMessage(v string) *EnableBackupLogResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *EnableBackupLogResponseBody) SetSuccess(v bool) *EnableBackupLogResponseBody {
	s.Success = &v
	return s
}

func (s *EnableBackupLogResponseBody) SetErrCode(v string) *EnableBackupLogResponseBody {
	s.ErrCode = &v
	return s
}

type EnableBackupLogResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableBackupLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableBackupLogResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupLogResponse) GoString() string {
	return s.String()
}

func (s *EnableBackupLogResponse) SetHeaders(v map[string]*string) *EnableBackupLogResponse {
	s.Headers = v
	return s
}

func (s *EnableBackupLogResponse) SetBody(v *EnableBackupLogResponseBody) *EnableBackupLogResponse {
	s.Body = v
	return s
}

type GetDBListFromAgentRequest struct {
	BackupGatewayId      *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s GetDBListFromAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDBListFromAgentRequest) GoString() string {
	return s.String()
}

func (s *GetDBListFromAgentRequest) SetBackupGatewayId(v int64) *GetDBListFromAgentRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *GetDBListFromAgentRequest) SetSourceEndpointRegion(v string) *GetDBListFromAgentRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *GetDBListFromAgentRequest) SetTaskId(v int64) *GetDBListFromAgentRequest {
	s.TaskId = &v
	return s
}

func (s *GetDBListFromAgentRequest) SetOwnerId(v string) *GetDBListFromAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDBListFromAgentRequest) SetClientToken(v string) *GetDBListFromAgentRequest {
	s.ClientToken = &v
	return s
}

type GetDBListFromAgentResponseBody struct {
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage     *string                               `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	DbList         *GetDBListFromAgentResponseBodyDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Struct"`
	ErrCode        *string                               `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s GetDBListFromAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDBListFromAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBListFromAgentResponseBody) SetRequestId(v string) *GetDBListFromAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetHttpStatusCode(v int32) *GetDBListFromAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetErrMessage(v string) *GetDBListFromAgentResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetSuccess(v bool) *GetDBListFromAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetDbList(v *GetDBListFromAgentResponseBodyDbList) *GetDBListFromAgentResponseBody {
	s.DbList = v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetErrCode(v string) *GetDBListFromAgentResponseBody {
	s.ErrCode = &v
	return s
}

type GetDBListFromAgentResponseBodyDbList struct {
	DbName []*string `json:"dbName,omitempty" xml:"dbName,omitempty" type:"Repeated"`
}

func (s GetDBListFromAgentResponseBodyDbList) String() string {
	return tea.Prettify(s)
}

func (s GetDBListFromAgentResponseBodyDbList) GoString() string {
	return s.String()
}

func (s *GetDBListFromAgentResponseBodyDbList) SetDbName(v []*string) *GetDBListFromAgentResponseBodyDbList {
	s.DbName = v
	return s
}

type GetDBListFromAgentResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDBListFromAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDBListFromAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDBListFromAgentResponse) GoString() string {
	return s.String()
}

func (s *GetDBListFromAgentResponse) SetHeaders(v map[string]*string) *GetDBListFromAgentResponse {
	s.Headers = v
	return s
}

func (s *GetDBListFromAgentResponse) SetBody(v *GetDBListFromAgentResponseBody) *GetDBListFromAgentResponse {
	s.Body = v
	return s
}

type ModifyBackupObjectsRequest struct {
	BackupPlanId  *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyBackupObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupObjectsRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupObjectsRequest) SetBackupPlanId(v string) *ModifyBackupObjectsRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupObjectsRequest) SetBackupObjects(v string) *ModifyBackupObjectsRequest {
	s.BackupObjects = &v
	return s
}

func (s *ModifyBackupObjectsRequest) SetOwnerId(v string) *ModifyBackupObjectsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupObjectsRequest) SetClientToken(v string) *ModifyBackupObjectsRequest {
	s.ClientToken = &v
	return s
}

type ModifyBackupObjectsResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	NeedPrecheck   *bool   `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ModifyBackupObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupObjectsResponseBody) SetRequestId(v string) *ModifyBackupObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetHttpStatusCode(v int32) *ModifyBackupObjectsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetBackupPlanId(v string) *ModifyBackupObjectsResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetNeedPrecheck(v bool) *ModifyBackupObjectsResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetErrMessage(v string) *ModifyBackupObjectsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetSuccess(v bool) *ModifyBackupObjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupObjectsResponseBody) SetErrCode(v string) *ModifyBackupObjectsResponseBody {
	s.ErrCode = &v
	return s
}

type ModifyBackupObjectsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupObjectsResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupObjectsResponse) SetHeaders(v map[string]*string) *ModifyBackupObjectsResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupObjectsResponse) SetBody(v *ModifyBackupObjectsResponseBody) *ModifyBackupObjectsResponse {
	s.Body = v
	return s
}

type ModifyBackupPlanNameRequest struct {
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupPlanName *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyBackupPlanNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanNameRequest) SetBackupPlanId(v string) *ModifyBackupPlanNameRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupPlanNameRequest) SetBackupPlanName(v string) *ModifyBackupPlanNameRequest {
	s.BackupPlanName = &v
	return s
}

func (s *ModifyBackupPlanNameRequest) SetOwnerId(v string) *ModifyBackupPlanNameRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPlanNameRequest) SetClientToken(v string) *ModifyBackupPlanNameRequest {
	s.ClientToken = &v
	return s
}

type ModifyBackupPlanNameResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ModifyBackupPlanNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanNameResponseBody) SetRequestId(v string) *ModifyBackupPlanNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetHttpStatusCode(v int32) *ModifyBackupPlanNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetBackupPlanId(v string) *ModifyBackupPlanNameResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetErrMessage(v string) *ModifyBackupPlanNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetSuccess(v bool) *ModifyBackupPlanNameResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupPlanNameResponseBody) SetErrCode(v string) *ModifyBackupPlanNameResponseBody {
	s.ErrCode = &v
	return s
}

type ModifyBackupPlanNameResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupPlanNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupPlanNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPlanNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanNameResponse) SetHeaders(v map[string]*string) *ModifyBackupPlanNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPlanNameResponse) SetBody(v *ModifyBackupPlanNameResponseBody) *ModifyBackupPlanNameResponse {
	s.Body = v
	return s
}

type ModifyBackupSetDownloadRulesRequest struct {
	BackupPlanId                        *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	OpenAutoDownload                    *bool   `json:"OpenAutoDownload,omitempty" xml:"OpenAutoDownload,omitempty"`
	BackupGatewayId                     *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	FullDataFormat                      *string `json:"FullDataFormat,omitempty" xml:"FullDataFormat,omitempty"`
	IncrementDataFormat                 *string `json:"IncrementDataFormat,omitempty" xml:"IncrementDataFormat,omitempty"`
	BackupSetDownloadDir                *string `json:"BackupSetDownloadDir,omitempty" xml:"BackupSetDownloadDir,omitempty"`
	BackupSetDownloadTargetType         *string `json:"BackupSetDownloadTargetType,omitempty" xml:"BackupSetDownloadTargetType,omitempty"`
	BackupSetDownloadTargetTypeLocation *string `json:"BackupSetDownloadTargetTypeLocation,omitempty" xml:"BackupSetDownloadTargetTypeLocation,omitempty"`
	OwnerId                             *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken                         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyBackupSetDownloadRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupSetDownloadRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupPlanId(v string) *ModifyBackupSetDownloadRulesRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetOpenAutoDownload(v bool) *ModifyBackupSetDownloadRulesRequest {
	s.OpenAutoDownload = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupGatewayId(v int64) *ModifyBackupSetDownloadRulesRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetFullDataFormat(v string) *ModifyBackupSetDownloadRulesRequest {
	s.FullDataFormat = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetIncrementDataFormat(v string) *ModifyBackupSetDownloadRulesRequest {
	s.IncrementDataFormat = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupSetDownloadDir(v string) *ModifyBackupSetDownloadRulesRequest {
	s.BackupSetDownloadDir = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupSetDownloadTargetType(v string) *ModifyBackupSetDownloadRulesRequest {
	s.BackupSetDownloadTargetType = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetBackupSetDownloadTargetTypeLocation(v string) *ModifyBackupSetDownloadRulesRequest {
	s.BackupSetDownloadTargetTypeLocation = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetOwnerId(v string) *ModifyBackupSetDownloadRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesRequest) SetClientToken(v string) *ModifyBackupSetDownloadRulesRequest {
	s.ClientToken = &v
	return s
}

type ModifyBackupSetDownloadRulesResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ModifyBackupSetDownloadRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupSetDownloadRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetRequestId(v string) *ModifyBackupSetDownloadRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetHttpStatusCode(v int32) *ModifyBackupSetDownloadRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetBackupPlanId(v string) *ModifyBackupSetDownloadRulesResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetErrMessage(v string) *ModifyBackupSetDownloadRulesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetSuccess(v bool) *ModifyBackupSetDownloadRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponseBody) SetErrCode(v string) *ModifyBackupSetDownloadRulesResponseBody {
	s.ErrCode = &v
	return s
}

type ModifyBackupSetDownloadRulesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupSetDownloadRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupSetDownloadRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupSetDownloadRulesResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetDownloadRulesResponse) SetHeaders(v map[string]*string) *ModifyBackupSetDownloadRulesResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponse) SetBody(v *ModifyBackupSetDownloadRulesResponseBody) *ModifyBackupSetDownloadRulesResponse {
	s.Body = v
	return s
}

type ModifyBackupSourceEndpointRequest struct {
	BackupPlanId               *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	SourceEndpointRegion       *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	SourceEndpointInstanceID   *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	SourceEndpointIP           *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	SourceEndpointPort         *int32  `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	SourceEndpointUserName     *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	SourceEndpointPassword     *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	BackupGatewayId            *int64  `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	BackupObjects              *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	SourceEndpointOracleSID    *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	CrossAliyunId              *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	CrossRoleName              *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	OwnerId                    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken                *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyBackupSourceEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupSourceEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupSourceEndpointRequest) SetBackupPlanId(v string) *ModifyBackupSourceEndpointRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointInstanceType(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointRegion(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointInstanceID(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointIP(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointPort(v int32) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointDatabaseName(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointUserName(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointPassword(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetBackupGatewayId(v int64) *ModifyBackupSourceEndpointRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetBackupObjects(v string) *ModifyBackupSourceEndpointRequest {
	s.BackupObjects = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetSourceEndpointOracleSID(v string) *ModifyBackupSourceEndpointRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetCrossAliyunId(v string) *ModifyBackupSourceEndpointRequest {
	s.CrossAliyunId = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetCrossRoleName(v string) *ModifyBackupSourceEndpointRequest {
	s.CrossRoleName = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetOwnerId(v string) *ModifyBackupSourceEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupSourceEndpointRequest) SetClientToken(v string) *ModifyBackupSourceEndpointRequest {
	s.ClientToken = &v
	return s
}

type ModifyBackupSourceEndpointResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	NeedPrecheck   *bool   `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ModifyBackupSourceEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupSourceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupSourceEndpointResponseBody) SetRequestId(v string) *ModifyBackupSourceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetHttpStatusCode(v int32) *ModifyBackupSourceEndpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetBackupPlanId(v string) *ModifyBackupSourceEndpointResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetNeedPrecheck(v bool) *ModifyBackupSourceEndpointResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetErrMessage(v string) *ModifyBackupSourceEndpointResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetSuccess(v bool) *ModifyBackupSourceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupSourceEndpointResponseBody) SetErrCode(v string) *ModifyBackupSourceEndpointResponseBody {
	s.ErrCode = &v
	return s
}

type ModifyBackupSourceEndpointResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupSourceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupSourceEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupSourceEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupSourceEndpointResponse) SetHeaders(v map[string]*string) *ModifyBackupSourceEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupSourceEndpointResponse) SetBody(v *ModifyBackupSourceEndpointResponseBody) *ModifyBackupSourceEndpointResponse {
	s.Body = v
	return s
}

type ModifyBackupStrategyRequest struct {
	BackupPlanId             *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupPeriod             *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupStartTime          *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStrategyType       *string `json:"BackupStrategyType,omitempty" xml:"BackupStrategyType,omitempty"`
	BackupLogIntervalSeconds *int32  `json:"BackupLogIntervalSeconds,omitempty" xml:"BackupLogIntervalSeconds,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyBackupStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupStrategyRequest) SetBackupPlanId(v string) *ModifyBackupStrategyRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetBackupPeriod(v string) *ModifyBackupStrategyRequest {
	s.BackupPeriod = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetBackupStartTime(v string) *ModifyBackupStrategyRequest {
	s.BackupStartTime = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetBackupStrategyType(v string) *ModifyBackupStrategyRequest {
	s.BackupStrategyType = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetBackupLogIntervalSeconds(v int32) *ModifyBackupStrategyRequest {
	s.BackupLogIntervalSeconds = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetOwnerId(v string) *ModifyBackupStrategyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupStrategyRequest) SetClientToken(v string) *ModifyBackupStrategyRequest {
	s.ClientToken = &v
	return s
}

type ModifyBackupStrategyResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	NeedPrecheck   *bool   `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ModifyBackupStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupStrategyResponseBody) SetRequestId(v string) *ModifyBackupStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetHttpStatusCode(v int32) *ModifyBackupStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetBackupPlanId(v string) *ModifyBackupStrategyResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetNeedPrecheck(v bool) *ModifyBackupStrategyResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetErrMessage(v string) *ModifyBackupStrategyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetSuccess(v bool) *ModifyBackupStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupStrategyResponseBody) SetErrCode(v string) *ModifyBackupStrategyResponseBody {
	s.ErrCode = &v
	return s
}

type ModifyBackupStrategyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupStrategyResponse) SetHeaders(v map[string]*string) *ModifyBackupStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupStrategyResponse) SetBody(v *ModifyBackupStrategyResponseBody) *ModifyBackupStrategyResponse {
	s.Body = v
	return s
}

type ModifyStorageStrategyRequest struct {
	BackupPlanId                      *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupRetentionPeriod             *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	DuplicationInfrequentAccessPeriod *int32  `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	DuplicationArchivePeriod          *int32  `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	OwnerId                           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken                       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyStorageStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyStorageStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyStorageStrategyRequest) SetBackupPlanId(v string) *ModifyStorageStrategyRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetBackupRetentionPeriod(v int32) *ModifyStorageStrategyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetDuplicationInfrequentAccessPeriod(v int32) *ModifyStorageStrategyRequest {
	s.DuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetDuplicationArchivePeriod(v int32) *ModifyStorageStrategyRequest {
	s.DuplicationArchivePeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetOwnerId(v string) *ModifyStorageStrategyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetClientToken(v string) *ModifyStorageStrategyRequest {
	s.ClientToken = &v
	return s
}

type ModifyStorageStrategyResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	NeedPrecheck   *bool   `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ModifyStorageStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyStorageStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStorageStrategyResponseBody) SetRequestId(v string) *ModifyStorageStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetHttpStatusCode(v int32) *ModifyStorageStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetBackupPlanId(v string) *ModifyStorageStrategyResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetNeedPrecheck(v bool) *ModifyStorageStrategyResponseBody {
	s.NeedPrecheck = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetErrMessage(v string) *ModifyStorageStrategyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetSuccess(v bool) *ModifyStorageStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyStorageStrategyResponseBody) SetErrCode(v string) *ModifyStorageStrategyResponseBody {
	s.ErrCode = &v
	return s
}

type ModifyStorageStrategyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyStorageStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyStorageStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyStorageStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyStorageStrategyResponse) SetHeaders(v map[string]*string) *ModifyStorageStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyStorageStrategyResponse) SetBody(v *ModifyStorageStrategyResponseBody) *ModifyStorageStrategyResponse {
	s.Body = v
	return s
}

type ReleaseBackupPlanRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReleaseBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *ReleaseBackupPlanRequest) SetBackupPlanId(v string) *ReleaseBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ReleaseBackupPlanRequest) SetOwnerId(v string) *ReleaseBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseBackupPlanRequest) SetClientToken(v string) *ReleaseBackupPlanRequest {
	s.ClientToken = &v
	return s
}

type ReleaseBackupPlanResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ReleaseBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseBackupPlanResponseBody) SetRequestId(v string) *ReleaseBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetHttpStatusCode(v int32) *ReleaseBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetBackupPlanId(v string) *ReleaseBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetErrMessage(v string) *ReleaseBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetSuccess(v bool) *ReleaseBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseBackupPlanResponseBody) SetErrCode(v string) *ReleaseBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

type ReleaseBackupPlanResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *ReleaseBackupPlanResponse) SetHeaders(v map[string]*string) *ReleaseBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *ReleaseBackupPlanResponse) SetBody(v *ReleaseBackupPlanResponseBody) *ReleaseBackupPlanResponse {
	s.Body = v
	return s
}

type RenewBackupPlanRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	Period       *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime     *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RenewBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *RenewBackupPlanRequest) SetBackupPlanId(v string) *RenewBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *RenewBackupPlanRequest) SetPeriod(v string) *RenewBackupPlanRequest {
	s.Period = &v
	return s
}

func (s *RenewBackupPlanRequest) SetUsedTime(v int32) *RenewBackupPlanRequest {
	s.UsedTime = &v
	return s
}

func (s *RenewBackupPlanRequest) SetOwnerId(v string) *RenewBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewBackupPlanRequest) SetClientToken(v string) *RenewBackupPlanRequest {
	s.ClientToken = &v
	return s
}

type RenewBackupPlanResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s RenewBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *RenewBackupPlanResponseBody) SetRequestId(v string) *RenewBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetHttpStatusCode(v int32) *RenewBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetBackupPlanId(v string) *RenewBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetErrMessage(v string) *RenewBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetOrderId(v string) *RenewBackupPlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetSuccess(v bool) *RenewBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *RenewBackupPlanResponseBody) SetErrCode(v string) *RenewBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

type RenewBackupPlanResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *RenewBackupPlanResponse) SetHeaders(v map[string]*string) *RenewBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *RenewBackupPlanResponse) SetBody(v *RenewBackupPlanResponseBody) *RenewBackupPlanResponse {
	s.Body = v
	return s
}

type StartBackupPlanRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s StartBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s StartBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *StartBackupPlanRequest) SetBackupPlanId(v string) *StartBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *StartBackupPlanRequest) SetOwnerId(v string) *StartBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *StartBackupPlanRequest) SetClientToken(v string) *StartBackupPlanRequest {
	s.ClientToken = &v
	return s
}

type StartBackupPlanResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s StartBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *StartBackupPlanResponseBody) SetRequestId(v string) *StartBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetHttpStatusCode(v int32) *StartBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetBackupPlanId(v string) *StartBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetErrMessage(v string) *StartBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetSuccess(v bool) *StartBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *StartBackupPlanResponseBody) SetErrCode(v string) *StartBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

type StartBackupPlanResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s StartBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *StartBackupPlanResponse) SetHeaders(v map[string]*string) *StartBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *StartBackupPlanResponse) SetBody(v *StartBackupPlanResponseBody) *StartBackupPlanResponse {
	s.Body = v
	return s
}

type StartRestoreTaskRequest struct {
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s StartRestoreTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreTaskRequest) GoString() string {
	return s.String()
}

func (s *StartRestoreTaskRequest) SetRestoreTaskId(v string) *StartRestoreTaskRequest {
	s.RestoreTaskId = &v
	return s
}

func (s *StartRestoreTaskRequest) SetOwnerId(v string) *StartRestoreTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRestoreTaskRequest) SetClientToken(v string) *StartRestoreTaskRequest {
	s.ClientToken = &v
	return s
}

type StartRestoreTaskResponseBody struct {
	RestoreTaskId  *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s StartRestoreTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartRestoreTaskResponseBody) SetRestoreTaskId(v string) *StartRestoreTaskResponseBody {
	s.RestoreTaskId = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetRequestId(v string) *StartRestoreTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetHttpStatusCode(v int32) *StartRestoreTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetErrMessage(v string) *StartRestoreTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetSuccess(v bool) *StartRestoreTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartRestoreTaskResponseBody) SetErrCode(v string) *StartRestoreTaskResponseBody {
	s.ErrCode = &v
	return s
}

type StartRestoreTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartRestoreTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartRestoreTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRestoreTaskResponse) GoString() string {
	return s.String()
}

func (s *StartRestoreTaskResponse) SetHeaders(v map[string]*string) *StartRestoreTaskResponse {
	s.Headers = v
	return s
}

func (s *StartRestoreTaskResponse) SetBody(v *StartRestoreTaskResponseBody) *StartRestoreTaskResponse {
	s.Body = v
	return s
}

type StartTaskRequest struct {
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s StartTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTaskRequest) GoString() string {
	return s.String()
}

func (s *StartTaskRequest) SetTaskId(v string) *StartTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartTaskRequest) SetOwnerId(v string) *StartTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTaskRequest) SetClientToken(v string) *StartTaskRequest {
	s.ClientToken = &v
	return s
}

type StartTaskResponseBody struct {
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	JobTypeName    *string `json:"JobTypeName,omitempty" xml:"JobTypeName,omitempty"`
}

func (s StartTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartTaskResponseBody) SetTaskId(v string) *StartTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartTaskResponseBody) SetRequestId(v string) *StartTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTaskResponseBody) SetHttpStatusCode(v int32) *StartTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartTaskResponseBody) SetErrMessage(v string) *StartTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartTaskResponseBody) SetSuccess(v bool) *StartTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartTaskResponseBody) SetErrCode(v string) *StartTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartTaskResponseBody) SetJobTypeName(v string) *StartTaskResponseBody {
	s.JobTypeName = &v
	return s
}

type StartTaskResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTaskResponse) GoString() string {
	return s.String()
}

func (s *StartTaskResponse) SetHeaders(v map[string]*string) *StartTaskResponse {
	s.Headers = v
	return s
}

func (s *StartTaskResponse) SetBody(v *StartTaskResponseBody) *StartTaskResponse {
	s.Body = v
	return s
}

type StopBackupPlanRequest struct {
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	StopMethod   *string `json:"StopMethod,omitempty" xml:"StopMethod,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s StopBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s StopBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *StopBackupPlanRequest) SetBackupPlanId(v string) *StopBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *StopBackupPlanRequest) SetStopMethod(v string) *StopBackupPlanRequest {
	s.StopMethod = &v
	return s
}

func (s *StopBackupPlanRequest) SetOwnerId(v string) *StopBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *StopBackupPlanRequest) SetClientToken(v string) *StopBackupPlanRequest {
	s.ClientToken = &v
	return s
}

type StopBackupPlanResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s StopBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *StopBackupPlanResponseBody) SetRequestId(v string) *StopBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetHttpStatusCode(v int32) *StopBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetBackupPlanId(v string) *StopBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetErrMessage(v string) *StopBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetSuccess(v bool) *StopBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *StopBackupPlanResponseBody) SetErrCode(v string) *StopBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

type StopBackupPlanResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s StopBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *StopBackupPlanResponse) SetHeaders(v map[string]*string) *StopBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *StopBackupPlanResponse) SetBody(v *StopBackupPlanResponseBody) *StopBackupPlanResponse {
	s.Body = v
	return s
}

type UpgradeBackupPlanRequest struct {
	BackupPlanId  *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpgradeBackupPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPlanRequest) SetBackupPlanId(v string) *UpgradeBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *UpgradeBackupPlanRequest) SetInstanceClass(v string) *UpgradeBackupPlanRequest {
	s.InstanceClass = &v
	return s
}

func (s *UpgradeBackupPlanRequest) SetOwnerId(v string) *UpgradeBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeBackupPlanRequest) SetClientToken(v string) *UpgradeBackupPlanRequest {
	s.ClientToken = &v
	return s
}

type UpgradeBackupPlanResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BackupPlanId   *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s UpgradeBackupPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPlanResponseBody) SetRequestId(v string) *UpgradeBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetHttpStatusCode(v int32) *UpgradeBackupPlanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetBackupPlanId(v string) *UpgradeBackupPlanResponseBody {
	s.BackupPlanId = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetErrMessage(v string) *UpgradeBackupPlanResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetOrderId(v string) *UpgradeBackupPlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetSuccess(v bool) *UpgradeBackupPlanResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeBackupPlanResponseBody) SetErrCode(v string) *UpgradeBackupPlanResponseBody {
	s.ErrCode = &v
	return s
}

type UpgradeBackupPlanResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeBackupPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPlanResponse) SetHeaders(v map[string]*string) *UpgradeBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *UpgradeBackupPlanResponse) SetBody(v *UpgradeBackupPlanResponseBody) *UpgradeBackupPlanResponse {
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

func (client *Client) ConfigureBackupPlanWithOptions(request *ConfigureBackupPlanRequest, runtime *util.RuntimeOptions) (_result *ConfigureBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureBackupPlan"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureBackupPlan(request *ConfigureBackupPlanRequest) (_result *ConfigureBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureBackupPlanResponse{}
	_body, _err := client.ConfigureBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupPlanWithOptions(request *CreateBackupPlanRequest, runtime *util.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBackupPlan"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackupPlan(request *CreateBackupPlanRequest) (_result *CreateBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CreateBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFullBackupSetDownloadWithOptions(request *CreateFullBackupSetDownloadRequest, runtime *util.RuntimeOptions) (_result *CreateFullBackupSetDownloadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFullBackupSetDownloadResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFullBackupSetDownload"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFullBackupSetDownload(request *CreateFullBackupSetDownloadRequest) (_result *CreateFullBackupSetDownloadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFullBackupSetDownloadResponse{}
	_body, _err := client.CreateFullBackupSetDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGetDBListFromAgentTaskWithOptions(request *CreateGetDBListFromAgentTaskRequest, runtime *util.RuntimeOptions) (_result *CreateGetDBListFromAgentTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGetDBListFromAgentTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGetDBListFromAgentTask"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGetDBListFromAgentTask(request *CreateGetDBListFromAgentTaskRequest) (_result *CreateGetDBListFromAgentTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGetDBListFromAgentTaskResponse{}
	_body, _err := client.CreateGetDBListFromAgentTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIncrementBackupSetDownloadWithOptions(request *CreateIncrementBackupSetDownloadRequest, runtime *util.RuntimeOptions) (_result *CreateIncrementBackupSetDownloadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIncrementBackupSetDownloadResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIncrementBackupSetDownload"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIncrementBackupSetDownload(request *CreateIncrementBackupSetDownloadRequest) (_result *CreateIncrementBackupSetDownloadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIncrementBackupSetDownloadResponse{}
	_body, _err := client.CreateIncrementBackupSetDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRestoreTaskWithOptions(request *CreateRestoreTaskRequest, runtime *util.RuntimeOptions) (_result *CreateRestoreTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRestoreTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRestoreTask"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRestoreTask(request *CreateRestoreTaskRequest) (_result *CreateRestoreTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRestoreTaskResponse{}
	_body, _err := client.CreateRestoreTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupGatewayListWithOptions(request *DescribeBackupGatewayListRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupGatewayListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupGatewayListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupGatewayList"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupGatewayList(request *DescribeBackupGatewayListRequest) (_result *DescribeBackupGatewayListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupGatewayListResponse{}
	_body, _err := client.DescribeBackupGatewayListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPlanBillingWithOptions(request *DescribeBackupPlanBillingRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPlanBillingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPlanBillingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPlanBilling"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPlanBilling(request *DescribeBackupPlanBillingRequest) (_result *DescribeBackupPlanBillingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPlanBillingResponse{}
	_body, _err := client.DescribeBackupPlanBillingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPlanListWithOptions(request *DescribeBackupPlanListRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPlanListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPlanListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPlanList"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPlanList(request *DescribeBackupPlanListRequest) (_result *DescribeBackupPlanListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPlanListResponse{}
	_body, _err := client.DescribeBackupPlanListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupSetDownloadTaskListWithOptions(request *DescribeBackupSetDownloadTaskListRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupSetDownloadTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupSetDownloadTaskListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupSetDownloadTaskList"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupSetDownloadTaskList(request *DescribeBackupSetDownloadTaskListRequest) (_result *DescribeBackupSetDownloadTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupSetDownloadTaskListResponse{}
	_body, _err := client.DescribeBackupSetDownloadTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFullBackupListWithOptions(request *DescribeFullBackupListRequest, runtime *util.RuntimeOptions) (_result *DescribeFullBackupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFullBackupListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFullBackupList"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFullBackupList(request *DescribeFullBackupListRequest) (_result *DescribeFullBackupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFullBackupListResponse{}
	_body, _err := client.DescribeFullBackupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIncrementBackupListWithOptions(request *DescribeIncrementBackupListRequest, runtime *util.RuntimeOptions) (_result *DescribeIncrementBackupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIncrementBackupListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIncrementBackupList"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIncrementBackupList(request *DescribeIncrementBackupListRequest) (_result *DescribeIncrementBackupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIncrementBackupListResponse{}
	_body, _err := client.DescribeIncrementBackupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobErrorCodeWithOptions(request *DescribeJobErrorCodeRequest, runtime *util.RuntimeOptions) (_result *DescribeJobErrorCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeJobErrorCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeJobErrorCode"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJobErrorCode(request *DescribeJobErrorCodeRequest) (_result *DescribeJobErrorCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobErrorCodeResponse{}
	_body, _err := client.DescribeJobErrorCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodeCidrListWithOptions(request *DescribeNodeCidrListRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeCidrListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNodeCidrListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNodeCidrList"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeCidrList(request *DescribeNodeCidrListRequest) (_result *DescribeNodeCidrListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeCidrListResponse{}
	_body, _err := client.DescribeNodeCidrListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePreCheckProgressListWithOptions(request *DescribePreCheckProgressListRequest, runtime *util.RuntimeOptions) (_result *DescribePreCheckProgressListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePreCheckProgressListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePreCheckProgressList"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePreCheckProgressList(request *DescribePreCheckProgressListRequest) (_result *DescribePreCheckProgressListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePreCheckProgressListResponse{}
	_body, _err := client.DescribePreCheckProgressListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreRangeInfoWithOptions(request *DescribeRestoreRangeInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreRangeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRestoreRangeInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRestoreRangeInfo"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreRangeInfo(request *DescribeRestoreRangeInfoRequest) (_result *DescribeRestoreRangeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreRangeInfoResponse{}
	_body, _err := client.DescribeRestoreRangeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRestoreTaskListWithOptions(request *DescribeRestoreTaskListRequest, runtime *util.RuntimeOptions) (_result *DescribeRestoreTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRestoreTaskListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRestoreTaskList"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRestoreTaskList(request *DescribeRestoreTaskListRequest) (_result *DescribeRestoreTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRestoreTaskListResponse{}
	_body, _err := client.DescribeRestoreTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableBackupLogWithOptions(request *DisableBackupLogRequest, runtime *util.RuntimeOptions) (_result *DisableBackupLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableBackupLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableBackupLog"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableBackupLog(request *DisableBackupLogRequest) (_result *DisableBackupLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableBackupLogResponse{}
	_body, _err := client.DisableBackupLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableBackupLogWithOptions(request *EnableBackupLogRequest, runtime *util.RuntimeOptions) (_result *EnableBackupLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableBackupLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableBackupLog"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableBackupLog(request *EnableBackupLogRequest) (_result *EnableBackupLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableBackupLogResponse{}
	_body, _err := client.EnableBackupLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDBListFromAgentWithOptions(request *GetDBListFromAgentRequest, runtime *util.RuntimeOptions) (_result *GetDBListFromAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDBListFromAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDBListFromAgent"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDBListFromAgent(request *GetDBListFromAgentRequest) (_result *GetDBListFromAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDBListFromAgentResponse{}
	_body, _err := client.GetDBListFromAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupObjectsWithOptions(request *ModifyBackupObjectsRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupObjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupObjects"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupObjects(request *ModifyBackupObjectsRequest) (_result *ModifyBackupObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupObjectsResponse{}
	_body, _err := client.ModifyBackupObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPlanNameWithOptions(request *ModifyBackupPlanNameRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPlanNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupPlanNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPlanName"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPlanName(request *ModifyBackupPlanNameRequest) (_result *ModifyBackupPlanNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPlanNameResponse{}
	_body, _err := client.ModifyBackupPlanNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupSetDownloadRulesWithOptions(request *ModifyBackupSetDownloadRulesRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupSetDownloadRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupSetDownloadRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupSetDownloadRules"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupSetDownloadRules(request *ModifyBackupSetDownloadRulesRequest) (_result *ModifyBackupSetDownloadRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupSetDownloadRulesResponse{}
	_body, _err := client.ModifyBackupSetDownloadRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupSourceEndpointWithOptions(request *ModifyBackupSourceEndpointRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupSourceEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupSourceEndpointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupSourceEndpoint"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupSourceEndpoint(request *ModifyBackupSourceEndpointRequest) (_result *ModifyBackupSourceEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupSourceEndpointResponse{}
	_body, _err := client.ModifyBackupSourceEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupStrategyWithOptions(request *ModifyBackupStrategyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupStrategy"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupStrategy(request *ModifyBackupStrategyRequest) (_result *ModifyBackupStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupStrategyResponse{}
	_body, _err := client.ModifyBackupStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyStorageStrategyWithOptions(request *ModifyStorageStrategyRequest, runtime *util.RuntimeOptions) (_result *ModifyStorageStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyStorageStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyStorageStrategy"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyStorageStrategy(request *ModifyStorageStrategyRequest) (_result *ModifyStorageStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyStorageStrategyResponse{}
	_body, _err := client.ModifyStorageStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseBackupPlanWithOptions(request *ReleaseBackupPlanRequest, runtime *util.RuntimeOptions) (_result *ReleaseBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseBackupPlan"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseBackupPlan(request *ReleaseBackupPlanRequest) (_result *ReleaseBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseBackupPlanResponse{}
	_body, _err := client.ReleaseBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewBackupPlanWithOptions(request *RenewBackupPlanRequest, runtime *util.RuntimeOptions) (_result *RenewBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewBackupPlan"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewBackupPlan(request *RenewBackupPlanRequest) (_result *RenewBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewBackupPlanResponse{}
	_body, _err := client.RenewBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartBackupPlanWithOptions(request *StartBackupPlanRequest, runtime *util.RuntimeOptions) (_result *StartBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartBackupPlan"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartBackupPlan(request *StartBackupPlanRequest) (_result *StartBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartBackupPlanResponse{}
	_body, _err := client.StartBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartRestoreTaskWithOptions(request *StartRestoreTaskRequest, runtime *util.RuntimeOptions) (_result *StartRestoreTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartRestoreTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartRestoreTask"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartRestoreTask(request *StartRestoreTaskRequest) (_result *StartRestoreTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRestoreTaskResponse{}
	_body, _err := client.StartRestoreTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTaskWithOptions(request *StartTaskRequest, runtime *util.RuntimeOptions) (_result *StartTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartTask"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTask(request *StartTaskRequest) (_result *StartTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTaskResponse{}
	_body, _err := client.StartTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopBackupPlanWithOptions(request *StopBackupPlanRequest, runtime *util.RuntimeOptions) (_result *StopBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopBackupPlan"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopBackupPlan(request *StopBackupPlanRequest) (_result *StopBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopBackupPlanResponse{}
	_body, _err := client.StopBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeBackupPlanWithOptions(request *UpgradeBackupPlanRequest, runtime *util.RuntimeOptions) (_result *UpgradeBackupPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeBackupPlanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeBackupPlan"), tea.String("2019-03-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeBackupPlan(request *UpgradeBackupPlanRequest) (_result *UpgradeBackupPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeBackupPlanResponse{}
	_body, _err := client.UpgradeBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
