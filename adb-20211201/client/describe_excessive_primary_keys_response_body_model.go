// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcessivePrimaryKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeExcessivePrimaryKeysResponseBody
	GetAccessDeniedDetail() *string
	SetDBClusterId(v string) *DescribeExcessivePrimaryKeysResponseBody
	GetDBClusterId() *string
	SetDetectionItems(v []*DescribeExcessivePrimaryKeysResponseBodyDetectionItems) *DescribeExcessivePrimaryKeysResponseBody
	GetDetectionItems() []*DescribeExcessivePrimaryKeysResponseBodyDetectionItems
	SetPageNumber(v int32) *DescribeExcessivePrimaryKeysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeExcessivePrimaryKeysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeExcessivePrimaryKeysResponseBody
	GetRequestId() *string
	SetTables(v []*DescribeExcessivePrimaryKeysResponseBodyTables) *DescribeExcessivePrimaryKeysResponseBody
	GetTables() []*DescribeExcessivePrimaryKeysResponseBodyTables
	SetTotalCount(v string) *DescribeExcessivePrimaryKeysResponseBody
	GetTotalCount() *string
}

type DescribeExcessivePrimaryKeysResponseBody struct {
	// The queried information about the request denial.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1906102576997697",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcNqPqHV6lwR4INiAGjIvK1ngXxN1O+6ORRB6A8YvztEOGywOk81ZmuNk0YrNy+qk7+UVDTHeXKsy8h9e/ePY/LMidj0RCmDpo/YpCumd0UGe0qEPe2U+UJAm/+UHlnEFLVg6BP3yIB5D++MCy7mgWm8Kwyhk62IeYly4hQ+5IpXjkh1GQXuDgLVVPVpxEek9n30vnCUL4KsaMgfa7dgojb+3TM8xGsD2zVK5STJNrsXclscIJEqyNXd7CBYiRJVZi1HPO6drN9WW0chLpCSTgjO8n0bNanZaxXKumW9PSwV58UoSFASeMWfZK3TLngX+oq8nGmnTwcJosVjfF4RGzAnS1IXt0Q9N2WHDnpwyLBU/nOz7Hsy8IZ+h+OVjsBTXSM9688/vOF707a5mNzpETvQeGRcua3A5livcKAM2cML0yeUs/Zyj/+BGqtVa+wektspDHC/CECh6R5lxQjRmUdPawY8VDs2onmdLuEH8DdmYt+Yv/jBFBUMWOyAluzkPYcX5nuQKouCIUJUFTSbsJsuH5CTIh7Ls5rbmkj+T1qTVz8gnDR8LxwaqoMSna+elXgVyOOxXtMkenVntsmoC3p/4G7yTPL1hu8JyWGIIvZHZGGLXGEH7FeSuMV8buKxPGFWG3arG8e9LGvDdz5dgTien4y6G5AQ0o1iQdXDos5VWdH3u7k5PrsvdEOpvMi6uSd8a42na80FsYlgGlwM5upydcWUC5Un2HCkJpT1xgk2L6shdVTrK6bidRrqE784FhW9bBQePzGaxSupPENZya0VUctRt+7uq3QwIn4y5jzjgX0E0jgmqPrgiVDjBesMQZYfGPCGysWYWYzfoh+G6V7N2VVGtNnGUwNWzM0WJBPONAgxPv+AmixFRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "202515810214480629",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeExcessivePrimaryKeysResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 863D51B7-5321-41D8-A0B6-A088B0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried tables that have excessive primary key fields.
	Tables []*DescribeExcessivePrimaryKeysResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 300
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExcessivePrimaryKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcessivePrimaryKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcessivePrimaryKeysResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeExcessivePrimaryKeysResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeExcessivePrimaryKeysResponseBody) GetDetectionItems() []*DescribeExcessivePrimaryKeysResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeExcessivePrimaryKeysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExcessivePrimaryKeysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExcessivePrimaryKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExcessivePrimaryKeysResponseBody) GetTables() []*DescribeExcessivePrimaryKeysResponseBodyTables {
	return s.Tables
}

func (s *DescribeExcessivePrimaryKeysResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeExcessivePrimaryKeysResponseBody) SetAccessDeniedDetail(v string) *DescribeExcessivePrimaryKeysResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBody) SetDBClusterId(v string) *DescribeExcessivePrimaryKeysResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBody) SetDetectionItems(v []*DescribeExcessivePrimaryKeysResponseBodyDetectionItems) *DescribeExcessivePrimaryKeysResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBody) SetPageNumber(v int32) *DescribeExcessivePrimaryKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBody) SetPageSize(v int32) *DescribeExcessivePrimaryKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBody) SetRequestId(v string) *DescribeExcessivePrimaryKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBody) SetTables(v []*DescribeExcessivePrimaryKeysResponseBodyTables) *DescribeExcessivePrimaryKeysResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBody) SetTotalCount(v string) *DescribeExcessivePrimaryKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExcessivePrimaryKeysResponseBodyDetectionItems struct {
	// The detection result.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// test
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

func (s DescribeExcessivePrimaryKeysResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcessivePrimaryKeysResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeExcessivePrimaryKeysResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeExcessivePrimaryKeysResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeExcessivePrimaryKeysResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeExcessivePrimaryKeysResponseBodyDetectionItems) SetMessage(v string) *DescribeExcessivePrimaryKeysResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyDetectionItems) SetName(v string) *DescribeExcessivePrimaryKeysResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyDetectionItems) SetStatus(v string) *DescribeExcessivePrimaryKeysResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}

type DescribeExcessivePrimaryKeysResponseBodyTables struct {
	// The total number of columns.
	//
	// example:
	//
	// 21
	ColumnCount *int32 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	// The queried primary key fields.
	//
	// example:
	//
	// 2
	PrimaryKeyColumns *string `json:"PrimaryKeyColumns,omitempty" xml:"PrimaryKeyColumns,omitempty"`
	// The number of primary key fields.
	//
	// example:
	//
	// 3
	PrimaryKeyCount *int32 `json:"PrimaryKeyCount,omitempty" xml:"PrimaryKeyCount,omitempty"`
	// The data size of primary key indexes. Unit: bytes.
	//
	// example:
	//
	// 222
	PrimaryKeyIndexSize *int64 `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The percentage of the table size. Unit: %.
	//
	// >  Formula: Table storage percentage = Total data size of a table/Total data size of the cluster × 100%.
	//
	// example:
	//
	// 23
	SpaceRatio *float64 `json:"SpaceRatio,omitempty" xml:"SpaceRatio,omitempty"`
	// The name of the table
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The cold data size. Unit: bytes.
	//
	// >  Formula: Cold data size = Data size of table records + Data size of regular indexes + Data size of primary key indexes + Data size of other data.
	//
	// example:
	//
	// 4
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeExcessivePrimaryKeysResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcessivePrimaryKeysResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) GetColumnCount() *int32 {
	return s.ColumnCount
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) GetPrimaryKeyColumns() *string {
	return s.PrimaryKeyColumns
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) GetPrimaryKeyCount() *int32 {
	return s.PrimaryKeyCount
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) GetPrimaryKeyIndexSize() *int64 {
	return s.PrimaryKeyIndexSize
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) GetSpaceRatio() *float64 {
	return s.SpaceRatio
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) GetTableName() *string {
	return s.TableName
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) SetColumnCount(v int32) *DescribeExcessivePrimaryKeysResponseBodyTables {
	s.ColumnCount = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) SetPrimaryKeyColumns(v string) *DescribeExcessivePrimaryKeysResponseBodyTables {
	s.PrimaryKeyColumns = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) SetPrimaryKeyCount(v int32) *DescribeExcessivePrimaryKeysResponseBodyTables {
	s.PrimaryKeyCount = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) SetPrimaryKeyIndexSize(v int64) *DescribeExcessivePrimaryKeysResponseBodyTables {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) SetSchemaName(v string) *DescribeExcessivePrimaryKeysResponseBodyTables {
	s.SchemaName = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) SetSpaceRatio(v float64) *DescribeExcessivePrimaryKeysResponseBodyTables {
	s.SpaceRatio = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) SetTableName(v string) *DescribeExcessivePrimaryKeysResponseBodyTables {
	s.TableName = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) SetTotalSize(v int64) *DescribeExcessivePrimaryKeysResponseBodyTables {
	s.TotalSize = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysResponseBodyTables) Validate() error {
	return dara.Validate(s)
}
