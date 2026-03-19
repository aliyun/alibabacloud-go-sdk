// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIncrementBackupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeIncrementBackupListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeIncrementBackupListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeIncrementBackupListResponseBody
	GetHttpStatusCode() *int32
	SetItems(v *DescribeIncrementBackupListResponseBodyItems) *DescribeIncrementBackupListResponseBody
	GetItems() *DescribeIncrementBackupListResponseBodyItems
	SetPageNum(v int32) *DescribeIncrementBackupListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeIncrementBackupListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIncrementBackupListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeIncrementBackupListResponseBody
	GetSuccess() *bool
	SetTotalElements(v int32) *DescribeIncrementBackupListResponseBody
	GetTotalElements() *int32
	SetTotalPages(v int32) *DescribeIncrementBackupListResponseBody
	GetTotalPages() *int32
}

type DescribeIncrementBackupListResponseBody struct {
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
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          *DescribeIncrementBackupListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5D52069-E8AA-5056-8C5C-654C3610****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of incremental backup tasks.
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

func (s DescribeIncrementBackupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIncrementBackupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeIncrementBackupListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeIncrementBackupListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeIncrementBackupListResponseBody) GetItems() *DescribeIncrementBackupListResponseBodyItems {
	return s.Items
}

func (s *DescribeIncrementBackupListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeIncrementBackupListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIncrementBackupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIncrementBackupListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeIncrementBackupListResponseBody) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *DescribeIncrementBackupListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeIncrementBackupListResponseBody) SetErrCode(v string) *DescribeIncrementBackupListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetErrMessage(v string) *DescribeIncrementBackupListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetHttpStatusCode(v int32) *DescribeIncrementBackupListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetItems(v *DescribeIncrementBackupListResponseBodyItems) *DescribeIncrementBackupListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetPageNum(v int32) *DescribeIncrementBackupListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetPageSize(v int32) *DescribeIncrementBackupListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetRequestId(v string) *DescribeIncrementBackupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBody) SetSuccess(v bool) *DescribeIncrementBackupListResponseBody {
	s.Success = &v
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

func (s *DescribeIncrementBackupListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIncrementBackupListResponseBodyItems struct {
	IncrementBackupFile []*DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile `json:"IncrementBackupFile,omitempty" xml:"IncrementBackupFile,omitempty" type:"Repeated"`
}

func (s DescribeIncrementBackupListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeIncrementBackupListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListResponseBodyItems) GetIncrementBackupFile() []*DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	return s.IncrementBackupFile
}

func (s *DescribeIncrementBackupListResponseBodyItems) SetIncrementBackupFile(v []*DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) *DescribeIncrementBackupListResponseBodyItems {
	s.IncrementBackupFile = v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItems) Validate() error {
	if s.IncrementBackupFile != nil {
		for _, item := range s.IncrementBackupFile {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile struct {
	BackupSetExpiredTime *int64  `json:"BackupSetExpiredTime,omitempty" xml:"BackupSetExpiredTime,omitempty"`
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSetJobId       *string `json:"BackupSetJobId,omitempty" xml:"BackupSetJobId,omitempty"`
	BackupSize           *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStatus         *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	EndTime              *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StorageMethod        *string `json:"StorageMethod,omitempty" xml:"StorageMethod,omitempty"`
}

func (s DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) String() string {
	return dara.Prettify(s)
}

func (s DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetBackupSetExpiredTime() *int64 {
	return s.BackupSetExpiredTime
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetBackupSetJobId() *string {
	return s.BackupSetJobId
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetSourceEndpointHost() *string {
	return s.SourceEndpointHost
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetSourceEndpointInstanceId() *string {
	return s.SourceEndpointInstanceId
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetSourceEndpointIpPort() *string {
	return s.SourceEndpointIpPort
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetSourceEndpointPort() *string {
	return s.SourceEndpointPort
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) GetStorageMethod() *string {
	return s.StorageMethod
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupSetExpiredTime(v int64) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupSetExpiredTime = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupSetId(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupSetId = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupSetJobId(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupSetJobId = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupSize(v int64) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupSize = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetBackupStatus(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.BackupStatus = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetEndTime(v int64) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.EndTime = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetSourceEndpointHost(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.SourceEndpointHost = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetSourceEndpointInstanceId(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.SourceEndpointInstanceId = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetSourceEndpointInstanceType(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetSourceEndpointIpPort(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.SourceEndpointIpPort = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetSourceEndpointPort(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.SourceEndpointPort = &v
	return s
}

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) SetSourceEndpointRegion(v string) *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile {
	s.SourceEndpointRegion = &v
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

func (s *DescribeIncrementBackupListResponseBodyItemsIncrementBackupFile) Validate() error {
	return dara.Validate(s)
}
