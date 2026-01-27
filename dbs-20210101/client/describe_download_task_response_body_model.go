// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDownloadTaskResponseBody
	GetCode() *string
	SetData(v *DescribeDownloadTaskResponseBodyData) *DescribeDownloadTaskResponseBody
	GetData() *DescribeDownloadTaskResponseBodyData
	SetErrCode(v string) *DescribeDownloadTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDownloadTaskResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeDownloadTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDownloadTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDownloadTaskResponseBody
	GetSuccess() *string
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
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDownloadTaskResponseBody) GetData() *DescribeDownloadTaskResponseBodyData {
	return s.Data
}

func (s *DescribeDownloadTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDownloadTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDownloadTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDownloadTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadTaskResponseBody) GetSuccess() *string {
	return s.Success
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

func (s *DescribeDownloadTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBodyData) GetContent() *DescribeDownloadTaskResponseBodyDataContent {
	return s.Content
}

func (s *DescribeDownloadTaskResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *DescribeDownloadTaskResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDownloadTaskResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDownloadTaskResponseBodyData) GetTotalElements() *int64 {
	return s.TotalElements
}

func (s *DescribeDownloadTaskResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
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

func (s *DescribeDownloadTaskResponseBodyData) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDownloadTaskResponseBodyDataContent struct {
	List []*DescribeDownloadTaskResponseBodyDataContentList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeDownloadTaskResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBodyDataContent) GetList() []*DescribeDownloadTaskResponseBodyDataContentList {
	return s.List
}

func (s *DescribeDownloadTaskResponseBodyDataContent) SetList(v []*DescribeDownloadTaskResponseBodyDataContentList) *DescribeDownloadTaskResponseBodyDataContent {
	s.List = v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContent) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskResponseBodyDataContentList) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetBackupSetTime() *string {
	return s.BackupSetTime
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetBakSetId() *string {
	return s.BakSetId
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetDbList() *string {
	return s.DbList
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetDownloadStatus() *string {
	return s.DownloadStatus
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetExportDataSize() *string {
	return s.ExportDataSize
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetFormat() *string {
	return s.Format
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetImportDataSize() *string {
	return s.ImportDataSize
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetTargetPath() *string {
	return s.TargetPath
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) GetTaskId() *string {
	return s.TaskId
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

func (s *DescribeDownloadTaskResponseBodyDataContentList) Validate() error {
	return dara.Validate(s)
}
