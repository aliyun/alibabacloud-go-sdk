// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataPushResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeDataPushResultRequest
	GetAppId() *string
	SetDataNames(v string) *DescribeDataPushResultRequest
	GetDataNames() *string
	SetDataVersions(v string) *DescribeDataPushResultRequest
	GetDataVersions() *string
	SetMaxDate(v string) *DescribeDataPushResultRequest
	GetMaxDate() *string
	SetMinDate(v string) *DescribeDataPushResultRequest
	GetMinDate() *string
	SetPageNumber(v int32) *DescribeDataPushResultRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataPushResultRequest
	GetPageSize() *int32
	SetRegionIds(v string) *DescribeDataPushResultRequest
	GetRegionIds() *string
}

type DescribeDataPushResultRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e76f8985-7965-41fc-925b-9648bb6bf650
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The names of the data files. This is a filter condition for the query. Multiple values are supported and separated by commas (,). By default, all data files are queried.
	//
	// example:
	//
	// 159828628258496/mirror_file/game-2553efe7-7bf8-40fb-a6e7-09c9c00a992a.tar
	DataNames *string `json:"DataNames,omitempty" xml:"DataNames,omitempty"`
	// The version numbers of the data files. This is a filter condition for the query. Multiple values are supported and separated by commas (,). By default, all versions are queried.
	//
	// example:
	//
	// 90396
	DataVersions *string `json:"DataVersions,omitempty" xml:"DataVersions,omitempty"`
	// The end time of the query. Format: YYYY-MM-DD. By default, there is no time limit.
	//
	// example:
	//
	// 2022-02-20
	MaxDate *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	// The start time of the query. Format: YYYY-MM-DD. By default, there is no time limit.
	//
	// example:
	//
	// 2022-02-15
	MinDate *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	// The page number. Pages start from 1. This parameter is not required when you export all information.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page set in Settings for paging during a paged query. This parameter is not required when you export all entries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of the edge nodes. This is a filter condition for the query. Multiple values are supported and separated by commas (,). By default, all edge nodes are queried.
	//
	// example:
	//
	// cn-wuhan-telecom_unicom_cmcc-2,cn-jiaozuo-2
	RegionIds *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
}

func (s DescribeDataPushResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeDataPushResultRequest) GetDataNames() *string {
	return s.DataNames
}

func (s *DescribeDataPushResultRequest) GetDataVersions() *string {
	return s.DataVersions
}

func (s *DescribeDataPushResultRequest) GetMaxDate() *string {
	return s.MaxDate
}

func (s *DescribeDataPushResultRequest) GetMinDate() *string {
	return s.MinDate
}

func (s *DescribeDataPushResultRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataPushResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataPushResultRequest) GetRegionIds() *string {
	return s.RegionIds
}

func (s *DescribeDataPushResultRequest) SetAppId(v string) *DescribeDataPushResultRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetDataNames(v string) *DescribeDataPushResultRequest {
	s.DataNames = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetDataVersions(v string) *DescribeDataPushResultRequest {
	s.DataVersions = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetMaxDate(v string) *DescribeDataPushResultRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetMinDate(v string) *DescribeDataPushResultRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetPageNumber(v int32) *DescribeDataPushResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetPageSize(v int32) *DescribeDataPushResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataPushResultRequest) SetRegionIds(v string) *DescribeDataPushResultRequest {
	s.RegionIds = &v
	return s
}

func (s *DescribeDataPushResultRequest) Validate() error {
	return dara.Validate(s)
}
