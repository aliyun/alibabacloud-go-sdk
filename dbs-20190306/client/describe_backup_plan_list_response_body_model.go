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
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Details of the backup plans.
	Items *DescribeBackupPlanListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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
	// Whether the operation was successful.
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
	// Backup gateway ID.
	//
	// example:
	//
	// 827362187368736
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// Backup method. The return values are as follows:
	//
	// - **logical**: Logical backup
	//
	// - **physical**: Physical backup
	//
	// - **duplication**: Replication backup
	//
	// example:
	//
	// logical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// Backup objects.
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// Full backup cycle. The return values are as follows:
	//
	// - **Monday**: Monday
	//
	// - **Tuesday**: Tuesday
	//
	// - **Wednesday**: Wednesday
	//
	// - **Thursday**: Thursday
	//
	// - **Friday**: Friday
	//
	// - **Saturday**: Saturday
	//
	// - **Sunday**: Sunday
	//
	// example:
	//
	// Monday
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// Timestamp of the backup plan creation.
	//
	// example:
	//
	// 1582527713000
	BackupPlanCreateTime *int64 `json:"BackupPlanCreateTime,omitempty" xml:"BackupPlanCreateTime,omitempty"`
	// Backup plan ID.
	//
	// example:
	//
	// dbstooi01eXXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// Backup plan name.
	//
	// example:
	//
	// dbstooi01e****
	BackupPlanName *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	// Backup plan status. The return values are as follows:
	//
	// - **wait**: Not configured
	//
	// - **init**: Not started (pre-check failed)
	//
	// - **running**: Running
	//
	// - **stop**: Failed
	//
	// - **pause**: Paused
	//
	// - **locked**: Locked
	//
	// - **check_pass**: Pre-check passed
	//
	// example:
	//
	// init
	BackupPlanStatus *string `json:"BackupPlanStatus,omitempty" xml:"BackupPlanStatus,omitempty"`
	// Backup data retention period, with a value range of 0 to 1825 days.
	//
	// example:
	//
	// 365
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Download server directory of the backup set
	BackupSetDownloadDir *string `json:"BackupSetDownloadDir,omitempty" xml:"BackupSetDownloadDir,omitempty"`
	// Full data format for backup set download:
	//
	// 	- **Native**
	//
	// 	- **SQL**
	//
	// 	- **CSV**
	//
	// 	- **JSON**
	//
	// example:
	//
	// SQL
	BackupSetDownloadFullDataFormat *string `json:"BackupSetDownloadFullDataFormat,omitempty" xml:"BackupSetDownloadFullDataFormat,omitempty"`
	// Backup set download backup gateway ID.
	//
	// example:
	//
	// 123123
	BackupSetDownloadGatewayId *int64 `json:"BackupSetDownloadGatewayId,omitempty" xml:"BackupSetDownloadGatewayId,omitempty"`
	// Backup set download full data format:
	//
	// 	- **Native**
	//
	// 	- **SQL**
	//
	// 	- **CSV**
	//
	// 	- **JSON**
	//
	// example:
	//
	// SQL
	BackupSetDownloadIncrementDataFormat *string `json:"BackupSetDownloadIncrementDataFormat,omitempty" xml:"BackupSetDownloadIncrementDataFormat,omitempty"`
	// Backup set download target type.
	//
	// > The only value is: agent, indicating that the backup gateway is installed.
	//
	// example:
	//
	// agent
	BackupSetDownloadTargetType *string `json:"BackupSetDownloadTargetType,omitempty" xml:"BackupSetDownloadTargetType,omitempty"`
	// Full backup start time, in the format HH:mm.
	//
	// example:
	//
	// 14:22
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// Built-in storage type. The return values are as follows:
	//
	// - Empty (default): Backup data is stored on the user\\"s OSS.
	//
	// - system: Backup data is stored on the built-in OSS of DBS.
	//
	// example:
	//
	// system
	BackupStorageType *string `json:"BackupStorageType,omitempty" xml:"BackupStorageType,omitempty"`
	// Start time for the database restore period, with a return value of 1554560477000.
	//
	// example:
	//
	// 1554560477000
	BeginTimestampForRestore *int64 `json:"BeginTimestampForRestore,omitempty" xml:"BeginTimestampForRestore,omitempty"`
	// UID for cross-Aliyun account backup.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// RAM role name for cross-Aliyun account backup.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// Database type.
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// Time (in days) to convert to archive cold backup storage.
	//
	// example:
	//
	// 365
	DuplicationArchivePeriod *int32 `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	// Time (in days) to convert to infrequent access storage.
	//
	// example:
	//
	// 180
	DuplicationInfrequentAccessPeriod *int32 `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	// Indicates whether incremental log backup is enabled, with return values as follows:
	//
	// - **true**: Enabled
	//
	// - **false**: Disabled
	//
	// example:
	//
	// true
	EnableBackupLog *bool `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// End time of the database restorable period, in timestamp format.
	//
	// example:
	//
	// 1554560477000
	EndTimestampForRestore *int64 `json:"EndTimestampForRestore,omitempty" xml:"EndTimestampForRestore,omitempty"`
	// Pre-check task error message.
	//
	// example:
	//
	// can not connect to oracle instance orcl with user dbs
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Instance class, with return values as follows:
	//
	// - **micro**: Entry-level
	//
	// - **small**: Low-spec
	//
	// - **medium**: Medium-spec
	//
	// - **large**: High-spec
	//
	// - **xlarge**: High-spec (no traffic limit)
	//
	// example:
	//
	// micro
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// OSS Bucket name.
	//
	// example:
	//
	// dbs-backup-1857XXXXX489
	OSSBucketName *string `json:"OSSBucketName,omitempty" xml:"OSSBucketName,omitempty"`
	// OSS Bucket region.
	//
	// example:
	//
	// cn-hangzhou
	OSSBucketRegion *string `json:"OSSBucketRegion,omitempty" xml:"OSSBucketRegion,omitempty"`
	// Indicates whether the automatic backup set download feature is enabled.
	//
	// example:
	//
	// true
	OpenBackupSetAutoDownload *bool `json:"OpenBackupSetAutoDownload,omitempty" xml:"OpenBackupSetAutoDownload,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aekzecovzti****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Database name.
	//
	// example:
	//
	// test
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// Database instance ID.
	//
	// example:
	//
	// test
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// Location of the database, the return values are as follows:
	//
	// - **rds**
	//
	// - **ecs**
	//
	// - **express**: Database connected via dedicated line/VPN gateway/smart gateway
	//
	// - **agent**: Database connected via backup gateway
	//
	// - **dds**: Cloud MongoDB
	//
	// - **other**: Database connected directly via IP:Port
	//
	// example:
	//
	// rds
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// Database connection address.
	//
	// example:
	//
	// 100.*.*.10:33204
	SourceEndpointIpPort *string `json:"SourceEndpointIpPort,omitempty" xml:"SourceEndpointIpPort,omitempty"`
	// Oracle SID name.
	//
	// example:
	//
	// test
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// Database region.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// Database username.
	//
	// example:
	//
	// test
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
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

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetEndTimestampForRestore() *int64 {
	return s.EndTimestampForRestore
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetInstanceClass() *string {
	return s.InstanceClass
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

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetBackupGatewayId(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.BackupGatewayId = &v
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

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetEndTimestampForRestore(v int64) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.EndTimestampForRestore = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetErrMessage(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetInstanceClass(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.InstanceClass = &v
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

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointRegion(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) SetSourceEndpointUserName(v string) *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail {
	s.SourceEndpointUserName = &v
	return s
}

func (s *DescribeBackupPlanListResponseBodyItemsBackupPlanDetail) Validate() error {
	return dara.Validate(s)
}
