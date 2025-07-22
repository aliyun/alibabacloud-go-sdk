// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceMissingIndexListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceMissingIndexListResponseBody
	GetCode() *string
	SetData(v *GetInstanceMissingIndexListResponseBodyData) *GetInstanceMissingIndexListResponseBody
	GetData() *GetInstanceMissingIndexListResponseBodyData
	SetMessage(v string) *GetInstanceMissingIndexListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceMissingIndexListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetInstanceMissingIndexListResponseBody
	GetSuccess() *string
}

type GetInstanceMissingIndexListResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *GetInstanceMissingIndexListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A74B755-98B7-59DB-8724-1321B394****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceMissingIndexListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMissingIndexListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceMissingIndexListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceMissingIndexListResponseBody) GetData() *GetInstanceMissingIndexListResponseBodyData {
	return s.Data
}

func (s *GetInstanceMissingIndexListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceMissingIndexListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceMissingIndexListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetInstanceMissingIndexListResponseBody) SetCode(v string) *GetInstanceMissingIndexListResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBody) SetData(v *GetInstanceMissingIndexListResponseBodyData) *GetInstanceMissingIndexListResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceMissingIndexListResponseBody) SetMessage(v string) *GetInstanceMissingIndexListResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBody) SetRequestId(v string) *GetInstanceMissingIndexListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBody) SetSuccess(v string) *GetInstanceMissingIndexListResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceMissingIndexListResponseBodyData struct {
	// The returned data.
	List []*GetInstanceMissingIndexListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 16
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInstanceMissingIndexListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMissingIndexListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceMissingIndexListResponseBodyData) GetList() []*GetInstanceMissingIndexListResponseBodyDataList {
	return s.List
}

func (s *GetInstanceMissingIndexListResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetInstanceMissingIndexListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetInstanceMissingIndexListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetInstanceMissingIndexListResponseBodyData) SetList(v []*GetInstanceMissingIndexListResponseBodyDataList) *GetInstanceMissingIndexListResponseBodyData {
	s.List = v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyData) SetPageNo(v int64) *GetInstanceMissingIndexListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyData) SetPageSize(v int64) *GetInstanceMissingIndexListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyData) SetTotal(v int64) *GetInstanceMissingIndexListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetInstanceMissingIndexListResponseBodyDataList struct {
	// The average cost savings.
	//
	// example:
	//
	// 4.67
	AvgTotalUserCost *float64 `json:"AvgTotalUserCost,omitempty" xml:"AvgTotalUserCost,omitempty"`
	// The performance improvement, in percentage.
	//
	// example:
	//
	// 98.3
	AvgUserImpact *float64 `json:"AvgUserImpact,omitempty" xml:"AvgUserImpact,omitempty"`
	// The statement used to create the missing indexes.
	//
	// example:
	//
	// CREATE INDEX [IX_CLOUDDBA_school_dbo_stu@col1_@col2] ON [school].[dbo].[stu]([col1],[col2],[col3]) INCLUDE ([col4],[col5]) WITH (FILLFACTOR = 90, ONLINE = OFF);
	CreateIndex *string `json:"CreateIndex,omitempty" xml:"CreateIndex,omitempty"`
	// The database name.
	//
	// example:
	//
	// school
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The index columns included in the equal operation.
	//
	// example:
	//
	// col1,col2,col3
	EqualityColumns *string `json:"EqualityColumns,omitempty" xml:"EqualityColumns,omitempty"`
	// The columns on which indexes are missing.
	//
	// example:
	//
	// col3,col4
	IncludedColumns *string `json:"IncludedColumns,omitempty" xml:"IncludedColumns,omitempty"`
	// The number of indexes.
	//
	// example:
	//
	// 1
	IndexCount *int64 `json:"IndexCount,omitempty" xml:"IndexCount,omitempty"`
	// The index columns included in the not equal operation.
	//
	// example:
	//
	// 2392
	InequalityColumns *string `json:"InequalityColumns,omitempty" xml:"InequalityColumns,omitempty"`
	// The last seek time of a user.
	//
	// example:
	//
	// 1702023327000
	LastUserSeek *int64 `json:"LastUserSeek,omitempty" xml:"LastUserSeek,omitempty"`
	// The object name.
	//
	// example:
	//
	// stu
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 5025
	ReservedPages *int64 `json:"ReservedPages,omitempty" xml:"ReservedPages,omitempty"`
	// The table size.
	//
	// example:
	//
	// 39.26
	ReservedSize *float64 `json:"ReservedSize,omitempty" xml:"ReservedSize,omitempty"`
	// The number of table rows.
	//
	// example:
	//
	// 226945
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The schema name.
	//
	// example:
	//
	// dbo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The number of scans.
	//
	// example:
	//
	// 0
	SystemScans *int64 `json:"SystemScans,omitempty" xml:"SystemScans,omitempty"`
	// The number of seeks.
	//
	// example:
	//
	// 0
	SystemSeeks *int64 `json:"SystemSeeks,omitempty" xml:"SystemSeeks,omitempty"`
	// The number of compilations.
	//
	// example:
	//
	// 2392
	UniqueCompiles *int64 `json:"UniqueCompiles,omitempty" xml:"UniqueCompiles,omitempty"`
	// The number of scans performed by users.
	//
	// example:
	//
	// 0
	UserScans *int64 `json:"UserScans,omitempty" xml:"UserScans,omitempty"`
	// The number of seeks performed by users.
	//
	// example:
	//
	// 1081
	UserSeeks *int64 `json:"UserSeeks,omitempty" xml:"UserSeeks,omitempty"`
}

func (s GetInstanceMissingIndexListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMissingIndexListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetAvgTotalUserCost() *float64 {
	return s.AvgTotalUserCost
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetAvgUserImpact() *float64 {
	return s.AvgUserImpact
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetCreateIndex() *string {
	return s.CreateIndex
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetEqualityColumns() *string {
	return s.EqualityColumns
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetIncludedColumns() *string {
	return s.IncludedColumns
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetIndexCount() *int64 {
	return s.IndexCount
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetInequalityColumns() *string {
	return s.InequalityColumns
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetLastUserSeek() *int64 {
	return s.LastUserSeek
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetReservedPages() *int64 {
	return s.ReservedPages
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetReservedSize() *float64 {
	return s.ReservedSize
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetRowCount() *int64 {
	return s.RowCount
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetSystemScans() *int64 {
	return s.SystemScans
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetSystemSeeks() *int64 {
	return s.SystemSeeks
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetUniqueCompiles() *int64 {
	return s.UniqueCompiles
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetUserScans() *int64 {
	return s.UserScans
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) GetUserSeeks() *int64 {
	return s.UserSeeks
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetAvgTotalUserCost(v float64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.AvgTotalUserCost = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetAvgUserImpact(v float64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.AvgUserImpact = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetCreateIndex(v string) *GetInstanceMissingIndexListResponseBodyDataList {
	s.CreateIndex = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetDatabaseName(v string) *GetInstanceMissingIndexListResponseBodyDataList {
	s.DatabaseName = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetEqualityColumns(v string) *GetInstanceMissingIndexListResponseBodyDataList {
	s.EqualityColumns = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetIncludedColumns(v string) *GetInstanceMissingIndexListResponseBodyDataList {
	s.IncludedColumns = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetIndexCount(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.IndexCount = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetInequalityColumns(v string) *GetInstanceMissingIndexListResponseBodyDataList {
	s.InequalityColumns = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetLastUserSeek(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.LastUserSeek = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetObjectName(v string) *GetInstanceMissingIndexListResponseBodyDataList {
	s.ObjectName = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetReservedPages(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.ReservedPages = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetReservedSize(v float64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.ReservedSize = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetRowCount(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.RowCount = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetSchemaName(v string) *GetInstanceMissingIndexListResponseBodyDataList {
	s.SchemaName = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetSystemScans(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.SystemScans = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetSystemSeeks(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.SystemSeeks = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetUniqueCompiles(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.UniqueCompiles = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetUserScans(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.UserScans = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) SetUserSeeks(v int64) *GetInstanceMissingIndexListResponseBodyDataList {
	s.UserSeeks = &v
	return s
}

func (s *GetInstanceMissingIndexListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
