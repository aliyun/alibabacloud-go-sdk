// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHiveTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorHiveTablesRequest
	GetClusterId() *string
	SetDateTime(v string) *ListDoctorHiveTablesRequest
	GetDateTime() *string
	SetMaxResults(v int32) *ListDoctorHiveTablesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHiveTablesRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorHiveTablesRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorHiveTablesRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorHiveTablesRequest
	GetRegionId() *string
	SetTableNames(v []*string) *ListDoctorHiveTablesRequest
	GetTableNames() []*string
}

type ListDoctorHiveTablesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specify the date in the ISO 8601 standard. For example, 2023-01-01 represents January 1, 2023.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The basis on which you want to sort the query results. Valid values:
	//
	// 	- partitionNum: the number of partitions.
	//
	// 	- totalFileCount: the total number of files.
	//
	// 	- largeFileCount: the number of large files. Large files are those with a size greater than 1 GB.
	//
	// 	- mediumFileCount: the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	//
	// 	- smallFileCount: the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	//
	// 	- tinyFileCount: the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	//
	// 	- emptyFileCount: the number of empty files. Empty files are those with a size of 0 MB.
	//
	// 	- largeFileRatio: the proportion of large files. Large files are those with a size greater than 1 GB.
	//
	// 	- mediumFileRatio: the proportion of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	//
	// 	- smallFileRatio: the proportion of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	//
	// 	- tinyFileRatio: the proportion of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	//
	// 	- emptyFileRatio: the proportion of empty files. Empty files are those with a size of 0 MB.
	//
	// 	- hotDataSize: the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	//
	// 	- WarmDataSize: the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	//
	// 	- coldDataSize: the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	//
	// 	- freezeDataSize: the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	//
	// 	- totalDataSize: the total amount of data.
	//
	// 	- hotDataRatio: the proportion of hot data. Hot data refers to data that is accessed in previous seven days.
	//
	// 	- WarmDataRatio: the proportion of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	//
	// 	- coldDataRatio: the proportion of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	//
	// 	- freezeDataRatio: the proportion of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	//
	// 	- totalFileDayGrowthCount: the daily increment of the total number of files.
	//
	// 	- largeFileDayGrowthCount: the daily increment of the number of large files. Large files are those with a size greater than 1 GB.
	//
	// 	- mediumFileDayGrowthCount: the daily increment of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	//
	// 	- smallFileDayGrowthCount: the daily increment of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	//
	// 	- tinyFileDayGrowthCount: the daily increment of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	//
	// 	- emptyFileDayGrowthCount: the daily increment of the number of empty files. Empty files are those with a size of 0 MB.
	//
	// 	- hotDataDayGrowthSize: the daily increment of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	//
	// 	- warmDataDayGrowthSize: the daily increment of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	//
	// 	- coldDataDayGrowthSize: the daily increment of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	//
	// 	- freezeDataDayGrowthSize: the daily increment of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	//
	// 	- totalDataDayGrowthSize: the daily increment of the amount of total data.
	//
	// 	- totalFileCountDayGrowthRatio: the day-to-day growth rate of the total number of files.
	//
	// 	- largeFileCountDayGrowthRatio: the day-to-day growth rate of the number of large files. Large files are those with a size greater than 1 GB.
	//
	// 	- mediumFileCountDayGrowthRatio: the day-to-day growth rate of the number of medium files. Medium files are those with a size greater than or equal to 128 MB and less than or equal to 1 GB.
	//
	// 	- smallFileCountDayGrowthRatio: the day-to-day growth rate of the number of small files. Small files are those with a size greater than or equal to 10 MB and less than 128 MB.
	//
	// 	- tinyFileCountDayGrowthRatio: the day-to-day growth rate of the number of very small files. Very small files are those with a size greater than 0 MB and less than 10 MB.
	//
	// 	- emptyFileCountDayGrowthRatio: the day-to-day growth rate of the number of empty files. Empty files are those with a size of 0 MB.
	//
	// 	- hotDataSizeDayGrowthRatio: the day-to-day growth rate of the amount of hot data. Hot data refers to data that is accessed in previous seven days.
	//
	// 	- warmDataSizeDayGrowthRatio: the day-to-day growth rate of the amount of warm data. Warm data refers to data that is not accessed for more than 7 days but is accessed in previous 30 days.
	//
	// 	- coldDataSizeDayGrowthRatio: the day-to-day growth rate of the amount of cold data. Cold data refers to data that is not accessed for more than 30 days but is accessed in previous 90 days.
	//
	// 	- freezeDataSizeDayGrowthRatio: the day-to-day growth rate of the amount of very cold data. Very cold data refers to data that is not accessed for more than 90 days.
	//
	// 	- totalDataSizeDayGrowthRatio: the day-to-day growth rate of the total amount of data.
	//
	// example:
	//
	// totalFileCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order in which you want to sort the query results. Valid value:
	//
	// 	- ASC: in ascending order
	//
	// 	- DESC: in descending order
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The table names, which are used to filter the query results.
	//
	// example:
	//
	// null
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s ListDoctorHiveTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorHiveTablesRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorHiveTablesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHiveTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHiveTablesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorHiveTablesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorHiveTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorHiveTablesRequest) GetTableNames() []*string {
	return s.TableNames
}

func (s *ListDoctorHiveTablesRequest) SetClusterId(v string) *ListDoctorHiveTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorHiveTablesRequest) SetDateTime(v string) *ListDoctorHiveTablesRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorHiveTablesRequest) SetMaxResults(v int32) *ListDoctorHiveTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHiveTablesRequest) SetNextToken(v string) *ListDoctorHiveTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHiveTablesRequest) SetOrderBy(v string) *ListDoctorHiveTablesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorHiveTablesRequest) SetOrderType(v string) *ListDoctorHiveTablesRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorHiveTablesRequest) SetRegionId(v string) *ListDoctorHiveTablesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorHiveTablesRequest) SetTableNames(v []*string) *ListDoctorHiveTablesRequest {
	s.TableNames = v
	return s
}

func (s *ListDoctorHiveTablesRequest) Validate() error {
	return dara.Validate(s)
}
