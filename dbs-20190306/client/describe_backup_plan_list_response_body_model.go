// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlanListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeBackupPlanListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBackupPlanListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeBackupPlanListResponseBody
	GetHttpStatusCode() *int32
	SetItems(v *DescribeBackupPlanListResponseBodyItems) *DescribeBackupPlanListResponseBody
	GetItems() *DescribeBackupPlanListResponseBodyItems
	SetPageNum(v int32) *DescribeBackupPlanListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeBackupPlanListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupPlanListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupPlanListResponseBody
	GetSuccess() *bool
	SetTotalElements(v int32) *DescribeBackupPlanListResponseBody
	GetTotalElements() *int32
	SetTotalPages(v int32) *DescribeBackupPlanListResponseBody
	GetTotalPages() *int32
}

type DescribeBackupPlanListResponseBody struct {
	// Error code.
	//
	// example:
	//
	// InvalidParameterValid
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          *DescribeBackupPlanListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of records per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of backup plans.
	//
	// example:
	//
	// 100
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 4
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeBackupPlanListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBackupPlanListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupPlanListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBackupPlanListResponseBody) GetItems() *DescribeBackupPlanListResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupPlanListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeBackupPlanListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupPlanListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPlanListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupPlanListResponseBody) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *DescribeBackupPlanListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeBackupPlanListResponseBody) SetErrCode(v string) *DescribeBackupPlanListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetErrMessage(v string) *DescribeBackupPlanListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetHttpStatusCode(v int32) *DescribeBackupPlanListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetItems(v *DescribeBackupPlanListResponseBodyItems) *DescribeBackupPlanListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetPageNum(v int32) *DescribeBackupPlanListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetPageSize(v int32) *DescribeBackupPlanListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetRequestId(v string) *DescribeBackupPlanListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBody) SetSuccess(v bool) *DescribeBackupPlanListResponseBody {
	s.Success = &v
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

func (s *DescribeBackupPlanListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPlanListResponseBodyItems struct {
	BackupPlanDetail []*DescribeBackupPlanListResponseBodyItemsBackupPlanDetail `json:"BackupPlanDetail,omitempty" xml:"BackupPlanDetail,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlanListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListResponseBodyItems) GetBackupPlanDetail() []*DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	return s.BackupPlanDetail
}

func (s *DescribeBackupPlanListResponseBodyItems) SetBackupPlanDetail(v []*DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) *DescribeBackupPlanListResponseBodyItems {
	s.BackupPlanDetail = v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItems) Validate() error {
	if s.BackupPlanDetail != nil {
		for _, item := range s.BackupPlanDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupPlanListResponseBodyItemsBackupPlanDetail struct {
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// example:
	//
	// TESTGATEWAY
	BackupGatewayIdentifier *string `json:"BackupGatewayIdentifier,omitempty" xml:"BackupGatewayIdentifier,omitempty"`
	BackupMethod            *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupObjects           *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	BackupPeriod            *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupPlanCreateTime    *int64  `json:"BackupPlanCreateTime,omitempty" xml:"BackupPlanCreateTime,omitempty"`
	BackupPlanId            *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupPlanName          *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	// example:
	//
	// cn-beijing
	BackupPlanRegion                     *string `json:"BackupPlanRegion,omitempty" xml:"BackupPlanRegion,omitempty"`
	BackupPlanStatus                     *string `json:"BackupPlanStatus,omitempty" xml:"BackupPlanStatus,omitempty"`
	BackupRetentionPeriod                *int32  `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	BackupSetDownloadDir                 *string `json:"BackupSetDownloadDir,omitempty" xml:"BackupSetDownloadDir,omitempty"`
	BackupSetDownloadFullDataFormat      *string `json:"BackupSetDownloadFullDataFormat,omitempty" xml:"BackupSetDownloadFullDataFormat,omitempty"`
	BackupSetDownloadGatewayId           *int64  `json:"BackupSetDownloadGatewayId,omitempty" xml:"BackupSetDownloadGatewayId,omitempty"`
	BackupSetDownloadIncrementDataFormat *string `json:"BackupSetDownloadIncrementDataFormat,omitempty" xml:"BackupSetDownloadIncrementDataFormat,omitempty"`
	BackupSetDownloadTargetType          *string `json:"BackupSetDownloadTargetType,omitempty" xml:"BackupSetDownloadTargetType,omitempty"`
	BackupStartTime                      *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStorageType                    *string `json:"BackupStorageType,omitempty" xml:"BackupStorageType,omitempty"`
	BeginTimestampForRestore             *int64  `json:"BeginTimestampForRestore,omitempty" xml:"BeginTimestampForRestore,omitempty"`
	CrossAliyunId                        *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	CrossRoleName                        *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	DatabaseType                         *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	DuplicationArchivePeriod             *int32  `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	DuplicationInfrequentAccessPeriod    *int32  `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	EnableBackupLog                      *bool   `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	EnableMysqlPhysicalBackupBinLog      *bool   `json:"EnableMysqlPhysicalBackupBinLog,omitempty" xml:"EnableMysqlPhysicalBackupBinLog,omitempty"`
	EndTimestampForRestore               *int64  `json:"EndTimestampForRestore,omitempty" xml:"EndTimestampForRestore,omitempty"`
	ErrMessage                           *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 365
	IncrementBackupRetentionPeriod *string `json:"IncrementBackupRetentionPeriod,omitempty" xml:"IncrementBackupRetentionPeriod,omitempty"`
	// example:
	//
	// 365
	IncrementDuplicationArchivePeriod *string `json:"IncrementDuplicationArchivePeriod,omitempty" xml:"IncrementDuplicationArchivePeriod,omitempty"`
	// example:
	//
	// 365
	IncrementDuplicationInfrequentAccessPeriod *string `json:"IncrementDuplicationInfrequentAccessPeriod,omitempty" xml:"IncrementDuplicationInfrequentAccessPeriod,omitempty"`
	// example:
	//
	// PREPAY
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceClass      *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// example:
	//
	// 1764051253000
	InstanceExpiredTimestamp *int64 `json:"InstanceExpiredTimestamp,omitempty" xml:"InstanceExpiredTimestamp,omitempty"`
	// example:
	//
	// 365
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// example:
	//
	// 365
	LogDuplicationArchivePeriod *string `json:"LogDuplicationArchivePeriod,omitempty" xml:"LogDuplicationArchivePeriod,omitempty"`
	// example:
	//
	// 365
	LogDuplicationInfrequentAccessPeriod *string `json:"LogDuplicationInfrequentAccessPeriod,omitempty" xml:"LogDuplicationInfrequentAccessPeriod,omitempty"`
	OSSBucketName                        *string `json:"OSSBucketName,omitempty" xml:"OSSBucketName,omitempty"`
	OSSBucketRegion                      *string `json:"OSSBucketRegion,omitempty" xml:"OSSBucketRegion,omitempty"`
	OpenBackupSetAutoDownload            *bool   `json:"OpenBackupSetAutoDownload,omitempty" xml:"OpenBackupSetAutoDownload,omitempty"`
	ResourceGroupId                      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SourceEndpointDatabaseName           *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// example:
	//
	// true
	SourceEndpointEnableSsl *string `json:"SourceEndpointEnableSsl,omitempty" xml:"SourceEndpointEnableSsl,omitempty"`
	// example:
	//
	// 127.0.0.1
	SourceEndpointHost         *string `json:"SourceEndpointHost,omitempty" xml:"SourceEndpointHost,omitempty"`
	SourceEndpointInstanceID   *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	SourceEndpointIpPort       *string `json:"SourceEndpointIpPort,omitempty" xml:"SourceEndpointIpPort,omitempty"`
	SourceEndpointOracleSID    *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// example:
	//
	// 3306
	SourceEndpointPort     *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	SourceEndpointRegion   *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// example:
	//
	// encrypted
	StorageEncryptMethod *string `json:"StorageEncryptMethod,omitempty" xml:"StorageEncryptMethod,omitempty"`
}

func (s DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupGatewayIdentifier() *string {
	return s.BackupGatewayIdentifier
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupObjects() *string {
	return s.BackupObjects
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupPeriod() *string {
	return s.BackupPeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupPlanCreateTime() *int64 {
	return s.BackupPlanCreateTime
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupPlanName() *string {
	return s.BackupPlanName
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupPlanRegion() *string {
	return s.BackupPlanRegion
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupPlanStatus() *string {
	return s.BackupPlanStatus
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupSetDownloadDir() *string {
	return s.BackupSetDownloadDir
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupSetDownloadFullDataFormat() *string {
	return s.BackupSetDownloadFullDataFormat
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupSetDownloadGatewayId() *int64 {
	return s.BackupSetDownloadGatewayId
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupSetDownloadIncrementDataFormat() *string {
	return s.BackupSetDownloadIncrementDataFormat
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupSetDownloadTargetType() *string {
	return s.BackupSetDownloadTargetType
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBackupStorageType() *string {
	return s.BackupStorageType
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetBeginTimestampForRestore() *int64 {
	return s.BeginTimestampForRestore
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetCrossAliyunId() *string {
	return s.CrossAliyunId
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetCrossRoleName() *string {
	return s.CrossRoleName
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetDuplicationArchivePeriod() *int32 {
	return s.DuplicationArchivePeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetDuplicationInfrequentAccessPeriod() *int32 {
	return s.DuplicationInfrequentAccessPeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetEnableBackupLog() *bool {
	return s.EnableBackupLog
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetEnableMysqlPhysicalBackupBinLog() *bool {
	return s.EnableMysqlPhysicalBackupBinLog
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetEndTimestampForRestore() *int64 {
	return s.EndTimestampForRestore
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetIncrementBackupRetentionPeriod() *string {
	return s.IncrementBackupRetentionPeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetIncrementDuplicationArchivePeriod() *string {
	return s.IncrementDuplicationArchivePeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetIncrementDuplicationInfrequentAccessPeriod() *string {
	return s.IncrementDuplicationInfrequentAccessPeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetInstanceExpiredTimestamp() *int64 {
	return s.InstanceExpiredTimestamp
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetLogBackupRetentionPeriod() *string {
	return s.LogBackupRetentionPeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetLogDuplicationArchivePeriod() *string {
	return s.LogDuplicationArchivePeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetLogDuplicationInfrequentAccessPeriod() *string {
	return s.LogDuplicationInfrequentAccessPeriod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetOSSBucketName() *string {
	return s.OSSBucketName
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetOSSBucketRegion() *string {
	return s.OSSBucketRegion
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetOpenBackupSetAutoDownload() *bool {
	return s.OpenBackupSetAutoDownload
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointDatabaseName() *string {
	return s.SourceEndpointDatabaseName
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointEnableSsl() *string {
	return s.SourceEndpointEnableSsl
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointHost() *string {
	return s.SourceEndpointHost
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointIpPort() *string {
	return s.SourceEndpointIpPort
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointOracleSID() *string {
	return s.SourceEndpointOracleSID
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointPort() *string {
	return s.SourceEndpointPort
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetStorageEncryptMethod() *string {
	return s.StorageEncryptMethod
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupGatewayId(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupGatewayId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupGatewayIdentifier(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupGatewayIdentifier = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupMethod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupObjects(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupObjects = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPeriod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanCreateTime(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanCreateTime = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanId(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanRegion(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanRegion = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupPlanStatus(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupPlanStatus = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupRetentionPeriod(v int32) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadDir(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadDir = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadFullDataFormat(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadFullDataFormat = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadGatewayId(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadGatewayId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadIncrementDataFormat(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadIncrementDataFormat = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupSetDownloadTargetType(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupSetDownloadTargetType = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupStartTime(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupStorageType(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupStorageType = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBeginTimestampForRestore(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BeginTimestampForRestore = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetCrossAliyunId(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.CrossAliyunId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetCrossRoleName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.CrossRoleName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetDatabaseType(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.DatabaseType = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetDuplicationArchivePeriod(v int32) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.DuplicationArchivePeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetDuplicationInfrequentAccessPeriod(v int32) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.DuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetEnableBackupLog(v bool) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetEnableMysqlPhysicalBackupBinLog(v bool) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.EnableMysqlPhysicalBackupBinLog = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetEndTimestampForRestore(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.EndTimestampForRestore = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetErrMessage(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetIncrementBackupRetentionPeriod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.IncrementBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetIncrementDuplicationArchivePeriod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.IncrementDuplicationArchivePeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetIncrementDuplicationInfrequentAccessPeriod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.IncrementDuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetInstanceChargeType(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetInstanceClass(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.InstanceClass = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetInstanceExpiredTimestamp(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.InstanceExpiredTimestamp = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetLogBackupRetentionPeriod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetLogDuplicationArchivePeriod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.LogDuplicationArchivePeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetLogDuplicationInfrequentAccessPeriod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.LogDuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetOSSBucketName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.OSSBucketName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetOSSBucketRegion(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.OSSBucketRegion = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetOpenBackupSetAutoDownload(v bool) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.OpenBackupSetAutoDownload = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetResourceGroupId(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointDatabaseName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointEnableSsl(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointEnableSsl = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointHost(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointHost = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointInstanceID(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointInstanceType(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointIpPort(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointIpPort = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointOracleSID(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointPort(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointPort = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointRegion(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointUserName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointUserName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetStorageEncryptMethod(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.StorageEncryptMethod = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) Validate() error {
	return dara.Validate(s)
}
