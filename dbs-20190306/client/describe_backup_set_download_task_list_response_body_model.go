// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetDownloadTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeBackupSetDownloadTaskListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBackupSetDownloadTaskListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeBackupSetDownloadTaskListResponseBody
	GetHttpStatusCode() *int32
	SetItems(v *DescribeBackupSetDownloadTaskListResponseBodyItems) *DescribeBackupSetDownloadTaskListResponseBody
	GetItems() *DescribeBackupSetDownloadTaskListResponseBodyItems
	SetPageNum(v int32) *DescribeBackupSetDownloadTaskListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeBackupSetDownloadTaskListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupSetDownloadTaskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupSetDownloadTaskListResponseBody
	GetSuccess() *bool
	SetTotalElements(v int32) *DescribeBackupSetDownloadTaskListResponseBody
	GetTotalElements() *int32
	SetTotalPages(v int32) *DescribeBackupSetDownloadTaskListResponseBody
	GetTotalPages() *int32
}

type DescribeBackupSetDownloadTaskListResponseBody struct {
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
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backup schedules.
	Items *DescribeBackupSetDownloadTaskListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6DC3D286-E0E6-5988-A558-2184984B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of backup set download tasks.
	//
	// example:
	//
	// 1
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeBackupSetDownloadTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetItems() *DescribeBackupSetDownloadTaskListResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetErrCode(v string) *DescribeBackupSetDownloadTaskListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetErrMessage(v string) *DescribeBackupSetDownloadTaskListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetHttpStatusCode(v int32) *DescribeBackupSetDownloadTaskListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetItems(v *DescribeBackupSetDownloadTaskListResponseBodyItems) *DescribeBackupSetDownloadTaskListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetPageNum(v int32) *DescribeBackupSetDownloadTaskListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetPageSize(v int32) *DescribeBackupSetDownloadTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetRequestId(v string) *DescribeBackupSetDownloadTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBody) SetSuccess(v bool) *DescribeBackupSetDownloadTaskListResponseBody {
	s.Success = &v
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

func (s *DescribeBackupSetDownloadTaskListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupSetDownloadTaskListResponseBodyItems struct {
	BackupSetDownloadTaskDetail []*DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail `json:"BackupSetDownloadTaskDetail,omitempty" xml:"BackupSetDownloadTaskDetail,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetDownloadTaskListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItems) GetBackupSetDownloadTaskDetail() []*DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	return s.BackupSetDownloadTaskDetail
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItems) SetBackupSetDownloadTaskDetail(v []*DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) *DescribeBackupSetDownloadTaskListResponseBodyItems {
	s.BackupSetDownloadTaskDetail = v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItems) Validate() error {
	if s.BackupSetDownloadTaskDetail != nil {
		for _, item := range s.BackupSetDownloadTaskDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail struct {
	// The backup gateway that is used to download the backup set. This parameter is available only if the value of the BackupSetDownloadWay parameter is auto.
	//
	// example:
	//
	// N/A
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The ID of the backup schedule.
	//
	// example:
	//
	// qhnuhyw3****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The code of the backup set.
	//
	// example:
	//
	// mysql-bin.00XXXX
	BackupSetCode *string `json:"BackupSetCode,omitempty" xml:"BackupSetCode,omitempty"`
	// The format in which the backup set is downloaded. Valid values:
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
	// Native
	BackupSetDataFormat *string `json:"BackupSetDataFormat,omitempty" xml:"BackupSetDataFormat,omitempty"`
	// The size of the data in the backup set.
	//
	// example:
	//
	// 237837
	BackupSetDataSize *int64 `json:"BackupSetDataSize,omitempty" xml:"BackupSetDataSize,omitempty"`
	// The type of the database.
	//
	// example:
	//
	// MySQL
	BackupSetDbType *string `json:"BackupSetDbType,omitempty" xml:"BackupSetDbType,omitempty"`
	// The timestamp that indicates when the backup set download task was created.
	//
	// example:
	//
	// 1642044001000
	BackupSetDownloadCreateTime *int64 `json:"BackupSetDownloadCreateTime,omitempty" xml:"BackupSetDownloadCreateTime,omitempty"`
	// The server directory to which the backup set is downloaded.
	//
	// > This parameter is available only if the value of the BackupSetDownloadWay parameter is auto.
	//
	// example:
	//
	// N/A
	BackupSetDownloadDir *string `json:"BackupSetDownloadDir,omitempty" xml:"BackupSetDownloadDir,omitempty"`
	// The timestamp that indicates when the backup set download task is complete.
	//
	// example:
	//
	// 1642044013000
	BackupSetDownloadFinishTime *int64 `json:"BackupSetDownloadFinishTime,omitempty" xml:"BackupSetDownloadFinishTime,omitempty"`
	// The public download URL of the backup set.
	//
	// > This parameter is available only if the value of the BackupSetDownloadWay parameter is manual and the conversion is complete.
	//
	// example:
	//
	// "//dbs-137XXXX-cn-hangzhou-1hr5cpbtmXXXX.oss-cn-hangzhou.example"
	BackupSetDownloadInternetUrl *string `json:"BackupSetDownloadInternetUrl,omitempty" xml:"BackupSetDownloadInternetUrl,omitempty"`
	// The internal download URL of the backup set.
	//
	// > This parameter is available only if the value of the BackupSetDownloadWay parameter is manual and the conversion is complete.
	//
	// example:
	//
	// "//dbs-13XXXX-cn-hangzhou-1hr5cpbtmXXXX.oss-cn-hangzhou-internal.example"
	BackupSetDownloadIntranetUrl *string `json:"BackupSetDownloadIntranetUrl,omitempty" xml:"BackupSetDownloadIntranetUrl,omitempty"`
	// The status of the backup set download task. Valid values:
	//
	// 	- **checking**: The backup set download task is being prechecked.
	//
	// 	- **init**: The backup set download task is not started and fails to pass the precheck.
	//
	// 	- **check_pass**: The backup set download task passes the precheck.
	//
	// 	- **pause**: The backup set download task is paused.
	//
	// 	- **schedule**: The backup set download task is waiting to be scheduled
	//
	// 	- **running**: The backup set download task is running.
	//
	// 	- **stop**: The backup set download task fails.
	//
	// 	- **finish**: The backup set download task is complete.
	//
	// example:
	//
	// finish
	BackupSetDownloadStatus *string `json:"BackupSetDownloadStatus,omitempty" xml:"BackupSetDownloadStatus,omitempty"`
	// The type of the destination server to which the backup set is downloaded.
	//
	// > This parameter is available only if the value of the BackupSetDownloadWay parameter is auto. A value of agent indicates a server on which a backup gateway is installed.
	//
	// example:
	//
	// N/A
	BackupSetDownloadTargetType *string `json:"BackupSetDownloadTargetType,omitempty" xml:"BackupSetDownloadTargetType,omitempty"`
	// The ID of the backup set download task.
	//
	// example:
	//
	// urxgrxt7****
	BackupSetDownloadTaskId *string `json:"BackupSetDownloadTaskId,omitempty" xml:"BackupSetDownloadTaskId,omitempty"`
	// The name of the backup set download task.
	//
	// example:
	//
	// 1h7za2yws****-ManualCont
	BackupSetDownloadTaskName *string `json:"BackupSetDownloadTaskName,omitempty" xml:"BackupSetDownloadTaskName,omitempty"`
	// The method in which the backup set is downloaded. Valid values:
	//
	// 	- **manual**: The backup set is manually downloaded.
	//
	// 	- **auto**: The backup set is automatically downloaded.
	//
	// example:
	//
	// manual
	BackupSetDownloadWay *string `json:"BackupSetDownloadWay,omitempty" xml:"BackupSetDownloadWay,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 1h7za2yws****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The type of the backup set task. Valid values:
	//
	// 	- **cbs_backup_sub_full**: logical full backup set download task
	//
	// 	- **cbs_backup_sub_cont**: logical incremental backup set download task
	//
	// example:
	//
	// cbs_backup_sub_cont
	BackupSetJobType *string `json:"BackupSetJobType,omitempty" xml:"BackupSetJobType,omitempty"`
	// The error message.
	//
	// example:
	//
	// java.lang.IndexOutOfBoundsException.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetCode() *string {
	return s.BackupSetCode
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDataFormat() *string {
	return s.BackupSetDataFormat
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDataSize() *int64 {
	return s.BackupSetDataSize
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDbType() *string {
	return s.BackupSetDbType
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadCreateTime() *int64 {
	return s.BackupSetDownloadCreateTime
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadDir() *string {
	return s.BackupSetDownloadDir
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadFinishTime() *int64 {
	return s.BackupSetDownloadFinishTime
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadInternetUrl() *string {
	return s.BackupSetDownloadInternetUrl
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadIntranetUrl() *string {
	return s.BackupSetDownloadIntranetUrl
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadStatus() *string {
	return s.BackupSetDownloadStatus
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadTargetType() *string {
	return s.BackupSetDownloadTargetType
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadTaskId() *string {
	return s.BackupSetDownloadTaskId
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadTaskName() *string {
	return s.BackupSetDownloadTaskName
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetDownloadWay() *string {
	return s.BackupSetDownloadWay
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetBackupSetJobType() *string {
	return s.BackupSetJobType
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupGatewayId(v int64) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupGatewayId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupPlanId(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetCode(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetCode = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDataFormat(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDataFormat = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDataSize(v int64) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDataSize = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDbType(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDbType = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadCreateTime(v int64) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadCreateTime = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadDir(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadDir = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadFinishTime(v int64) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadFinishTime = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadInternetUrl(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadInternetUrl = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadIntranetUrl(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadIntranetUrl = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadStatus(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadStatus = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadTargetType(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadTargetType = &v
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

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetDownloadWay(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetDownloadWay = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetId(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetBackupSetJobType(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.BackupSetJobType = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) SetErrMessage(v string) *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupSetDownloadTaskListResponseBodyItemsBackupSetDownloadTaskDetail) Validate() error {
	return dara.Validate(s)
}
