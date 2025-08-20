// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInclinedTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeInclinedTablesResponseBody
	GetAccessDeniedDetail() *string
	SetDetectionItems(v []*DescribeInclinedTablesResponseBodyDetectionItems) *DescribeInclinedTablesResponseBody
	GetDetectionItems() []*DescribeInclinedTablesResponseBodyDetectionItems
	SetItems(v *DescribeInclinedTablesResponseBodyItems) *DescribeInclinedTablesResponseBody
	GetItems() *DescribeInclinedTablesResponseBodyItems
	SetPageNumber(v string) *DescribeInclinedTablesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeInclinedTablesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeInclinedTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeInclinedTablesResponseBody
	GetTotalCount() *string
}

type DescribeInclinedTablesResponseBody struct {
	// The details about the access denial. This parameter is returned only if Resource Access Management (RAM) permission verification failed.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeInclinedTablesResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The queried tables.
	Items *DescribeInclinedTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInclinedTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeInclinedTablesResponseBody) GetDetectionItems() []*DescribeInclinedTablesResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeInclinedTablesResponseBody) GetItems() *DescribeInclinedTablesResponseBodyItems {
	return s.Items
}

func (s *DescribeInclinedTablesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeInclinedTablesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInclinedTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInclinedTablesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeInclinedTablesResponseBody) SetAccessDeniedDetail(v string) *DescribeInclinedTablesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetDetectionItems(v []*DescribeInclinedTablesResponseBodyDetectionItems) *DescribeInclinedTablesResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetItems(v *DescribeInclinedTablesResponseBodyItems) *DescribeInclinedTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetPageNumber(v string) *DescribeInclinedTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetPageSize(v string) *DescribeInclinedTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetRequestId(v string) *DescribeInclinedTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) SetTotalCount(v string) *DescribeInclinedTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInclinedTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInclinedTablesResponseBodyDetectionItems struct {
	// The message of the detection result.
	//
	// example:
	//
	// A shard skew is detected in a table.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// Table skew
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The severity level of the detection result.
	//
	// example:
	//
	// NORMAL
	//
	// WARNNING
	//
	// CRITICAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInclinedTablesResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) SetMessage(v string) *DescribeInclinedTablesResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) SetName(v string) *DescribeInclinedTablesResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) SetStatus(v string) *DescribeInclinedTablesResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}

type DescribeInclinedTablesResponseBodyItems struct {
	// The queried table.
	Table []*DescribeInclinedTablesResponseBodyItemsTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeInclinedTablesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItems) GetTable() []*DescribeInclinedTablesResponseBodyItemsTable {
	return s.Table
}

func (s *DescribeInclinedTablesResponseBodyItems) SetTable(v []*DescribeInclinedTablesResponseBodyItemsTable) *DescribeInclinedTablesResponseBodyItems {
	s.Table = v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeInclinedTablesResponseBodyItemsTable struct {
	// Indicates whether data is skewed in the table.
	//
	// example:
	//
	// True
	IsIncline *bool `json:"IsIncline,omitempty" xml:"IsIncline,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// admin_daily_own_statistic_record
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 1000
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// spark_test
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 200
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The percentage of the table size. Unit: %.
	//
	// >  Formula: Table storage percentage = Total data size of a table/Total data size of the cluster × 100%.
	//
	// example:
	//
	// 0.4
	SpaceRatio *float64 `json:"SpaceRatio,omitempty" xml:"SpaceRatio,omitempty"`
	// The total data size of the table. Unit: bytes.
	//
	// >  The following formulas can be used to calculate the total data size:
	//
	// 	- Formula 1: Total data size = Hot data size + Cold data size.
	//
	// 	- Formula 2: Total data size = Data size of table records + Data size of regular indexes + Data size of primary key indexes + Data size of other data.
	//
	// example:
	//
	// 53687091200
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// The detection type of the table.
	//
	// example:
	//
	// Fact
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInclinedTablesResponseBodyItemsTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetIsIncline() *bool {
	return s.IsIncline
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetName() *string {
	return s.Name
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetRowCount() *int64 {
	return s.RowCount
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetSchema() *string {
	return s.Schema
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetSize() *int64 {
	return s.Size
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetSpaceRatio() *float64 {
	return s.SpaceRatio
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) GetType() *string {
	return s.Type
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetIsIncline(v bool) *DescribeInclinedTablesResponseBodyItemsTable {
	s.IsIncline = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetName(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Name = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetRowCount(v int64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.RowCount = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSchema(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Schema = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSize(v int64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Size = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetSpaceRatio(v float64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.SpaceRatio = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetTotalSize(v int64) *DescribeInclinedTablesResponseBodyItemsTable {
	s.TotalSize = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) SetType(v string) *DescribeInclinedTablesResponseBodyItemsTable {
	s.Type = &v
	return s
}

func (s *DescribeInclinedTablesResponseBodyItemsTable) Validate() error {
	return dara.Validate(s)
}
