// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRecordList(v []*ListBackupRecordResponseBodyBackupRecordList) *ListBackupRecordResponseBody
	GetBackupRecordList() []*ListBackupRecordResponseBodyBackupRecordList
	SetPageInfo(v *ListBackupRecordResponseBodyPageInfo) *ListBackupRecordResponseBody
	GetPageInfo() *ListBackupRecordResponseBodyPageInfo
	SetRequestId(v string) *ListBackupRecordResponseBody
	GetRequestId() *string
}

type ListBackupRecordResponseBody struct {
	// The details of the backup record.
	BackupRecordList []*ListBackupRecordResponseBodyBackupRecordList `json:"BackupRecordList,omitempty" xml:"BackupRecordList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListBackupRecordResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB48****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBackupRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBackupRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackupRecordResponseBody) GetBackupRecordList() []*ListBackupRecordResponseBodyBackupRecordList {
	return s.BackupRecordList
}

func (s *ListBackupRecordResponseBody) GetPageInfo() *ListBackupRecordResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListBackupRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBackupRecordResponseBody) SetBackupRecordList(v []*ListBackupRecordResponseBodyBackupRecordList) *ListBackupRecordResponseBody {
	s.BackupRecordList = v
	return s
}

func (s *ListBackupRecordResponseBody) SetPageInfo(v *ListBackupRecordResponseBodyPageInfo) *ListBackupRecordResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListBackupRecordResponseBody) SetRequestId(v string) *ListBackupRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBackupRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBackupRecordResponseBodyBackupRecordList struct {
	// The timestamp when the backup task ended. Unit: milliseconds.
	//
	// example:
	//
	// 1699600611000
	BackupEndTime *int64 `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The backup task ID.
	//
	// example:
	//
	// a006f24d069843c88688672d74ee****
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The backup plan ID.
	//
	// example:
	//
	// plan-000c4tt43nolmx96****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The timestamp when the backup task started. Unit: milliseconds.
	//
	// example:
	//
	// 1699514211000
	BackupStartTime *int64 `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The backup task status. Valid value:
	//
	// 	- **BACKUP_COMPLETE**: The backup task is successful.
	//
	// 	- **BACKUP_FAILED**: The backup task failed.
	//
	// 	- **PARTIAL_COMPLETE**: The backup task is partially successful.
	//
	// example:
	//
	// BACKUP_COMPLETE
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The ID of the anti-ransomware agent.
	//
	// example:
	//
	// c-0002bgagelj3d2sc****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The error code of the backup task.
	//
	// example:
	//
	// FILE_CACHE_NO_SPACE
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message of the backup task.
	//
	// example:
	//
	// FILE_CACHE_NO_SPACE
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// i-wz9ikn44p46krnic****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the asset.
	//
	// example:
	//
	// openapi
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 1.1.1.1
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.1.1
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 1.1.1.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the region in which the backup is stored.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UUID of the server whose data is backed up based on the anti-ransomware policy.
	//
	// example:
	//
	// b93cccb9-f19f-4886-97fe-47df26ba****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListBackupRecordResponseBodyBackupRecordList) String() string {
	return dara.Prettify(s)
}

func (s ListBackupRecordResponseBodyBackupRecordList) GoString() string {
	return s.String()
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetBackupEndTime() *int64 {
	return s.BackupEndTime
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetBackupStartTime() *int64 {
	return s.BackupStartTime
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetClientId() *string {
	return s.ClientId
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetIp() *string {
	return s.Ip
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBackupRecordResponseBodyBackupRecordList) GetUuid() *string {
	return s.Uuid
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetBackupEndTime(v int64) *ListBackupRecordResponseBodyBackupRecordList {
	s.BackupEndTime = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetBackupJobId(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.BackupJobId = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetBackupPlanId(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.BackupPlanId = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetBackupStartTime(v int64) *ListBackupRecordResponseBodyBackupRecordList {
	s.BackupStartTime = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetBackupStatus(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.BackupStatus = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetClientId(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.ClientId = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetErrorCode(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.ErrorCode = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetErrorMessage(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.ErrorMessage = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetInstanceId(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.InstanceId = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetInstanceName(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.InstanceName = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetInternetIp(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.InternetIp = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetIntranetIp(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.IntranetIp = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetIp(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.Ip = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetRegionId(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.RegionId = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) SetUuid(v string) *ListBackupRecordResponseBodyBackupRecordList {
	s.Uuid = &v
	return s
}

func (s *ListBackupRecordResponseBodyBackupRecordList) Validate() error {
	return dara.Validate(s)
}

type ListBackupRecordResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBackupRecordResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListBackupRecordResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListBackupRecordResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListBackupRecordResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListBackupRecordResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBackupRecordResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBackupRecordResponseBodyPageInfo) SetCount(v int32) *ListBackupRecordResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListBackupRecordResponseBodyPageInfo) SetCurrentPage(v int32) *ListBackupRecordResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListBackupRecordResponseBodyPageInfo) SetPageSize(v int32) *ListBackupRecordResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListBackupRecordResponseBodyPageInfo) SetTotalCount(v int32) *ListBackupRecordResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListBackupRecordResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
