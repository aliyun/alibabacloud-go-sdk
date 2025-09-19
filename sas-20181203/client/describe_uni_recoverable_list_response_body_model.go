// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniRecoverableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeUniRecoverableListResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeUniRecoverableListResponseBody
	GetCurrentPage() *int32
	SetDatabase(v string) *DescribeUniRecoverableListResponseBody
	GetDatabase() *string
	SetPageSize(v int32) *DescribeUniRecoverableListResponseBody
	GetPageSize() *int32
	SetRecoverableInfoList(v []*DescribeUniRecoverableListResponseBodyRecoverableInfoList) *DescribeUniRecoverableListResponseBody
	GetRecoverableInfoList() []*DescribeUniRecoverableListResponseBodyRecoverableInfoList
	SetRequestId(v string) *DescribeUniRecoverableListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeUniRecoverableListResponseBody
	GetTotalCount() *int32
}

type DescribeUniRecoverableListResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// msdb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// An array that consists of the backup snapshots.
	RecoverableInfoList []*DescribeUniRecoverableListResponseBodyRecoverableInfoList `json:"RecoverableInfoList,omitempty" xml:"RecoverableInfoList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D0760****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUniRecoverableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniRecoverableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUniRecoverableListResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeUniRecoverableListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeUniRecoverableListResponseBody) GetDatabase() *string {
	return s.Database
}

func (s *DescribeUniRecoverableListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUniRecoverableListResponseBody) GetRecoverableInfoList() []*DescribeUniRecoverableListResponseBodyRecoverableInfoList {
	return s.RecoverableInfoList
}

func (s *DescribeUniRecoverableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUniRecoverableListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUniRecoverableListResponseBody) SetCount(v int32) *DescribeUniRecoverableListResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBody) SetCurrentPage(v int32) *DescribeUniRecoverableListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBody) SetDatabase(v string) *DescribeUniRecoverableListResponseBody {
	s.Database = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBody) SetPageSize(v int32) *DescribeUniRecoverableListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBody) SetRecoverableInfoList(v []*DescribeUniRecoverableListResponseBodyRecoverableInfoList) *DescribeUniRecoverableListResponseBody {
	s.RecoverableInfoList = v
	return s
}

func (s *DescribeUniRecoverableListResponseBody) SetRequestId(v string) *DescribeUniRecoverableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBody) SetTotalCount(v int32) *DescribeUniRecoverableListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUniRecoverableListResponseBodyRecoverableInfoList struct {
	// The timestamp of the first backup. Unit: milliseconds.
	//
	// example:
	//
	// 1671468180000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// The timestamp of the last backup. Unit: milliseconds.
	//
	// example:
	//
	// 1671468180000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The identifier of the point in time for restoration in the backup version that is used. The database is an Oracle database.
	//
	// example:
	//
	// 4529940.0
	ResetScn *string `json:"ResetScn,omitempty" xml:"ResetScn,omitempty"`
	// The point in time for restoration in the backup version that is used. The database is an Oracle database.
	//
	// example:
	//
	// 2021-01-30 08:04:36
	ResetTime *int64 `json:"ResetTime,omitempty" xml:"ResetTime,omitempty"`
	// The information about the database. This parameter is available when the database is a Microsoft SQL Server (MSSQL) database. The value is a JSON string. Valid values:
	//
	// 	- **name**: the name of the database
	//
	// 	- **files**: the path to the database files
	//
	// example:
	//
	// {
	//
	//       "files": {
	//
	//             "qtc": "F:\\\\database\\\\qtc.mdf",
	//
	//             "qtc_log": "F:\\\\database\\\\qtc_0.ldf"
	//
	//       },
	//
	//       "name": "qtc"
	//
	// }
	RestoreInfo *string `json:"RestoreInfo,omitempty" xml:"RestoreInfo,omitempty"`
}

func (s DescribeUniRecoverableListResponseBodyRecoverableInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniRecoverableListResponseBodyRecoverableInfoList) GoString() string {
	return s.String()
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) GetResetScn() *string {
	return s.ResetScn
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) GetResetTime() *int64 {
	return s.ResetTime
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) GetRestoreInfo() *string {
	return s.RestoreInfo
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) SetFirstTime(v int64) *DescribeUniRecoverableListResponseBodyRecoverableInfoList {
	s.FirstTime = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) SetLastTime(v int64) *DescribeUniRecoverableListResponseBodyRecoverableInfoList {
	s.LastTime = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) SetResetScn(v string) *DescribeUniRecoverableListResponseBodyRecoverableInfoList {
	s.ResetScn = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) SetResetTime(v int64) *DescribeUniRecoverableListResponseBodyRecoverableInfoList {
	s.ResetTime = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) SetRestoreInfo(v string) *DescribeUniRecoverableListResponseBodyRecoverableInfoList {
	s.RestoreInfo = &v
	return s
}

func (s *DescribeUniRecoverableListResponseBodyRecoverableInfoList) Validate() error {
	return dara.Validate(s)
}
