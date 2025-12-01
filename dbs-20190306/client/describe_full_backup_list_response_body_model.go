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
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          *DescribeFullBackupListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 636BC118-6080-4119-A6B5-C199CEC1037D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
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
	// [{\\"DBName\\":\\"test\\"}]
	BackupObjects        *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	BackupSetExpiredTime *int64  `json:"BackupSetExpiredTime,omitempty" xml:"BackupSetExpiredTime,omitempty"`
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSize           *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// example:
	//
	// finish
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// NULL
	ErrMessage           *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	FinishTime           *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	SourceEndpointIpPort *string `json:"SourceEndpointIpPort,omitempty" xml:"SourceEndpointIpPort,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Standard
	StorageMethod *string `json:"StorageMethod,omitempty" xml:"StorageMethod,omitempty"`
}

func (s DescribeFullBackupListResponseBodyItemsFullBackupFile) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullBackupListResponseBodyItemsFullBackupFile) GoString() string {
	return s.String()
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

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetSourceEndpointIpPort() *string {
	return s.SourceEndpointIpPort
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) GetStorageMethod() *string {
	return s.StorageMethod
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

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) SetSourceEndpointIpPort(v string) *DescribeFullBackupListResponseBodyItemsFullBackupFile {
	s.SourceEndpointIpPort = &v
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

func (s *DescribeFullBackupListResponseBodyItemsFullBackupFile) Validate() error {
	return dara.Validate(s)
}
