// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFullBackupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeFullBackupListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeFullBackupListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeFullBackupListResponseBody
	GetHttpStatusCode() *int32
	SetItems(v *DescribeFullBackupListResponseBodyItems) *DescribeFullBackupListResponseBody
	GetItems() *DescribeFullBackupListResponseBodyItems
	SetPageNum(v int32) *DescribeFullBackupListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeFullBackupListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFullBackupListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeFullBackupListResponseBody
	GetSuccess() *bool
	SetTotalElements(v int32) *DescribeFullBackupListResponseBody
	GetTotalElements() *int32
	SetTotalPages(v int32) *DescribeFullBackupListResponseBody
	GetTotalPages() *int32
}

type DescribeFullBackupListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          *DescribeFullBackupListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 636BC118-6080-4119-A6B5-C199CEC1037D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - **true**: The operation succeeded.
	//
	// - **false**: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of full backup jobs.
	//
	// example:
	//
	// 1
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeFullBackupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullBackupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeFullBackupListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeFullBackupListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeFullBackupListResponseBody) GetItems() *DescribeFullBackupListResponseBodyItems {
	return s.Items
}

func (s *DescribeFullBackupListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeFullBackupListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFullBackupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFullBackupListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFullBackupListResponseBody) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *DescribeFullBackupListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeFullBackupListResponseBody) SetErrCode(v string) *DescribeFullBackupListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetErrMessage(v string) *DescribeFullBackupListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetHttpStatusCode(v int32) *DescribeFullBackupListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetItems(v *DescribeFullBackupListResponseBodyItems) *DescribeFullBackupListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetPageNum(v int32) *DescribeFullBackupListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetPageSize(v int32) *DescribeFullBackupListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetRequestId(v string) *DescribeFullBackupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFullBackupListResponseBody) SetSuccess(v bool) *DescribeFullBackupListResponseBody {
	s.Success = &v
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

func (s *DescribeFullBackupListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFullBackupListResponseBodyItems struct {
	FullBackupFile []*DescribeFullBackupListResponseBodyItemsFullBackupFile `json:"FullBackupFile,omitempty" xml:"FullBackupFile,omitempty" type:"Repeated"`
}

func (s DescribeFullBackupListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullBackupListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListResponseBodyItems) GetFullBackupFile() []*DescribeFullBackupListResponseBodyItemsFullBackupFile {
	return s.FullBackupFile
}

func (s *DescribeFullBackupListResponseBodyItems) SetFullBackupFile(v []*DescribeFullBackupListResponseBodyItemsFullBackupFile) *DescribeFullBackupListResponseBodyItems {
	s.FullBackupFile = v
	return s
}

func (s *DescribeFullBackupListResponseBodyItems) Validate() error {
	if s.FullBackupFile != nil {
		for _, item := range s.FullBackupFile {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFullBackupListResponseBodyItemsFullBackupFile struct {
	// example:
	//
	// TESTGATEWAY
	BackupGatewayIdentifier *string `json:"BackupGatewayIdentifier,omitempty" xml:"BackupGatewayIdentifier,omitempty"`
	BackupObjects           *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	BackupSetExpiredTime    *int64  `json:"BackupSetExpiredTime,omitempty" xml:"BackupSetExpiredTime,omitempty"`
	BackupSetId             *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSize              *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStatus            *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	CreateTime              *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// example:
	//
	// ram-for-dbs
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrMessage    *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	FinishTime    *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 100
	LogicalFullBackupProgress *int32 `json:"LogicalFullBackupProgress,omitempty" xml:"LogicalFullBackupProgress,omitempty"`
	// example:
	//
	// 100
	LogicalStructureBackupProgress *int32 `json:"LogicalStructureBackupProgress,omitempty" xml:"LogicalStructureBackupProgress,omitempty"`
	// example:
	//
	// true
	SourceEndpointEnableSsl *string `json:"SourceEndpointEnableSsl,omitempty" xml:"SourceEndpointEnableSsl,omitempty"`
	// example:
	//
	// 127.0.0.1
	SourceEndpointHost *string `json:"SourceEndpointHost,omitempty" xml:"SourceEndpointHost,omitempty"`
	// example:
	//
	// rm-testxx
	SourceEndpointInstanceId *string `json:"SourceEndpointInstanceId,omitempty" xml:"SourceEndpointInstanceId,omitempty"`
	// example:
	//
	// rds
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	SourceEndpointIpPort       *string `json:"SourceEndpointIpPort,omitempty" xml:"SourceEndpointIpPort,omitempty"`
	// example:
	//
	// 3306
	SourceEndpointPort *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// example:
	//
	// cn-beijing
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// example:
	//
	// dbs_backup
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	StartTime              *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// encrypted
	StorageEncryptMethod *string `json:"StorageEncryptMethod,omitempty" xml:"StorageEncryptMethod,omitempty"`
	StorageMethod        *string `json:"StorageMethod,omitempty" xml:"StorageMethod,omitempty"`
}

func (s DescribeFullBackupListResponseBodyItemsFullBackupFile) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullBackupListResponseBodyItemsFullBackupFile) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetBackupGatewayIdentifier() *string {
	return s.BackupGatewayIdentifier
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetBackupObjects() *string {
	return s.BackupObjects
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetBackupSetExpiredTime() *int64 {
	return s.BackupSetExpiredTime
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetCrossAliyunId() *string {
	return s.CrossAliyunId
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetCrossRoleName() *string {
	return s.CrossRoleName
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetLogicalFullBackupProgress() *int32 {
	return s.LogicalFullBackupProgress
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetLogicalStructureBackupProgress() *int32 {
	return s.LogicalStructureBackupProgress
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointEnableSsl() *string {
	return s.SourceEndpointEnableSsl
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointHost() *string {
	return s.SourceEndpointHost
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointInstanceId() *string {
	return s.SourceEndpointInstanceId
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointIpPort() *string {
	return s.SourceEndpointIpPort
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointPort() *string {
	return s.SourceEndpointPort
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetStorageEncryptMethod() *string {
	return s.StorageEncryptMethod
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetStorageMethod() *string {
	return s.StorageMethod
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupGatewayIdentifier(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupGatewayIdentifier = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupObjects(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupObjects = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupSetExpiredTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupSetExpiredTime = &v
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

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetBackupStatus(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.BackupStatus = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetCreateTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.CreateTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetCrossAliyunId(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.CrossAliyunId = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetCrossRoleName(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.CrossRoleName = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetEndTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.EndTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetErrMessage(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.ErrMessage = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetFinishTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.FinishTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetLogicalFullBackupProgress(v int32) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.LogicalFullBackupProgress = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetLogicalStructureBackupProgress(v int32) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.LogicalStructureBackupProgress = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointEnableSsl(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointEnableSsl = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointHost(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointHost = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointInstanceId(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointInstanceId = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointInstanceType(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointIpPort(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointIpPort = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointPort(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointPort = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointRegion(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointUserName(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointUserName = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetStartTime(v int64) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.StartTime = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetStorageEncryptMethod(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.StorageEncryptMethod = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetStorageMethod(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.StorageMethod = &v
	return s
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) Validate() error {
	return dara.Validate(s)
}
