// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHDFSDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorHDFSDirectoriesRequest
	GetClusterId() *string
	SetDateTime(v string) *ListDoctorHDFSDirectoriesRequest
	GetDateTime() *string
	SetDirPath(v string) *ListDoctorHDFSDirectoriesRequest
	GetDirPath() *string
	SetMaxResults(v int32) *ListDoctorHDFSDirectoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHDFSDirectoriesRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorHDFSDirectoriesRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorHDFSDirectoriesRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorHDFSDirectoriesRequest
	GetRegionId() *string
}

type ListDoctorHDFSDirectoriesRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The date. The value is in the ISO 8601 format. For example, 2023-01-01 represents January 1, 2023.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The path of the directory. The directory depth cannot exceed five levels. If you do not specify this parameter, the analysis results of all directories are returned.
	//
	// example:
	//
	// /tmp/test
	DirPath *string `json:"DirPath,omitempty" xml:"DirPath,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The property based on which to sort the query results. Valid values:
	//
	// - totalFileCount: the total number of files.
	//
	// - largeFileCount: the number of large files. A large file is 1 GB or larger.
	//
	// - mediumFileCount: the number of medium-sized files. A medium-sized file is larger than 128 MB and smaller than 1 GB.
	//
	// - smallFileCount: the number of small files. A small file is larger than 10 MB and less than or equal to 128 MB.
	//
	// - tinyFileCount: the number of tiny files. A tiny file is larger than 0 MB and less than or equal to 10 MB.
	//
	// - emptyFileCount: the number of empty files. An empty file is 0 MB in size.
	//
	// - hotDataSize: the size of hot data. Hot data is data that was accessed in the last 7 days.
	//
	// - warmDataSize: the size of warm data. Warm data is data that was not accessed in the last 7 days but was accessed in the last 30 days.
	//
	// - coldDataSize: the size of cold data. Cold data is data that was not accessed in the last 30 days but was accessed in the last 90 days.
	//
	// - freezeDataSize: the size of freeze data. Freeze data is data that was not accessed in the last 90 days.
	//
	// - totalDataSize: the total data size.
	//
	// - totalFileDayGrowthCount: the daily increase in the total number of files.
	//
	// - largeFileDayGrowthCount: the daily increase in the number of large files. A large file is 1 GB or larger.
	//
	// - mediumFileDayGrowthCount: the daily increase in the number of medium-sized files. A medium-sized file is larger than 128 MB and smaller than 1 GB.
	//
	// - smallFileDayGrowthCount: the daily increase in the number of small files. A small file is larger than 10 MB and less than or equal to 128 MB.
	//
	// - tinyFileDayGrowthCount: the daily increase in the number of tiny files. A tiny file is larger than 0 MB and less than or equal to 10 MB.
	//
	// - emptyFileDayGrowthCount: the daily increase in the number of empty files. An empty file is 0 MB in size.
	//
	// - hotDataDayGrowthSize: the daily increase in the size of hot data. Hot data is data that was accessed in the last 7 days.
	//
	// - warmDataDayGrowthSize: the daily increase in the size of warm data. Warm data is data that was not accessed in the last 7 days but was accessed in the last 30 days.
	//
	// - coldDataDayGrowthSize: the daily increase in the size of cold data. Cold data is data that was not accessed in the last 30 days but was accessed in the last 90 days.
	//
	// - freezeDataDayGrowthSize: the daily increase in the size of freeze data. Freeze data is data that was not accessed in the last 90 days.
	//
	// - totalDataDayGrowthSize: the daily increase in the total data size.
	//
	// - totalFileCountDayGrowthRatio: the day-over-day growth ratio of the total number of files.
	//
	// - largeFileCountDayGrowthRatio: the day-over-day growth ratio of the number of large files. A large file is 1 GB or larger.
	//
	// - mediumFileCountDayGrowthRatio: the day-over-day growth ratio of the number of medium-sized files. A medium-sized file is larger than 128 MB and smaller than 1 GB.
	//
	// - smallFileCountDayGrowthRatio: the day-over-day growth ratio of the number of small files. A small file is larger than 10 MB and less than or equal to 128 MB.
	//
	// - tinyFileCountDayGrowthRatio: the day-over-day growth ratio of the number of tiny files. A tiny file is larger than 0 MB and less than or equal to 10 MB.
	//
	// - emptyFileCountDayGrowthRatio: the day-over-day growth ratio of the number of empty files. An empty file is 0 MB in size.
	//
	// example:
	//
	// smallFileCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The sort order. Valid values:
	//
	// - ASC: ascending
	//
	// - DESC: descending
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDoctorHDFSDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorHDFSDirectoriesRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorHDFSDirectoriesRequest) GetDirPath() *string {
	return s.DirPath
}

func (s *ListDoctorHDFSDirectoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHDFSDirectoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHDFSDirectoriesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorHDFSDirectoriesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorHDFSDirectoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorHDFSDirectoriesRequest) SetClusterId(v string) *ListDoctorHDFSDirectoriesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetDateTime(v string) *ListDoctorHDFSDirectoriesRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetDirPath(v string) *ListDoctorHDFSDirectoriesRequest {
	s.DirPath = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetMaxResults(v int32) *ListDoctorHDFSDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetNextToken(v string) *ListDoctorHDFSDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetOrderBy(v string) *ListDoctorHDFSDirectoriesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetOrderType(v string) *ListDoctorHDFSDirectoriesRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) SetRegionId(v string) *ListDoctorHDFSDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
