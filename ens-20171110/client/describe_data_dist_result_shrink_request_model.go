// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataDistResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeDataDistResultShrinkRequest
	GetAppId() *string
	SetDataNames(v string) *DescribeDataDistResultShrinkRequest
	GetDataNames() *string
	SetDataVersions(v string) *DescribeDataDistResultShrinkRequest
	GetDataVersions() *string
	SetEnsRegionIdsShrink(v string) *DescribeDataDistResultShrinkRequest
	GetEnsRegionIdsShrink() *string
	SetInstanceIds(v string) *DescribeDataDistResultShrinkRequest
	GetInstanceIds() *string
	SetMaxDate(v string) *DescribeDataDistResultShrinkRequest
	GetMaxDate() *string
	SetMinDate(v string) *DescribeDataDistResultShrinkRequest
	GetMinDate() *string
	SetPageNumber(v int32) *DescribeDataDistResultShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataDistResultShrinkRequest
	GetPageSize() *int32
}

type DescribeDataDistResultShrinkRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// e76f8985-7965-41fc-925b-9648bb6bf650
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the data file. Separate multiple names with commas (,). By default, all data files are queried.
	//
	// example:
	//
	// cloudgame-shanghai/deploy_app/20220215/1644895216305ACG_M21B-ota-1.1.2-D-0215.0628_V1_0002-pre-weiduan.zip
	DataNames *string `json:"DataNames,omitempty" xml:"DataNames,omitempty"`
	// The version number of the data file. Separate multiple numbers with commas (,). By default, all versions of data files are queried.
	//
	// example:
	//
	// 4885
	DataVersions *string `json:"DataVersions,omitempty" xml:"DataVersions,omitempty"`
	// The IDs of the ENS nodes.
	EnsRegionIdsShrink *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	// The IDs of ENS instances. Separate multiple IDs with commas (,). By default, all instances are queried.
	//
	// example:
	//
	// i-7ecpqvkicnchxccozrp,i-6ecpqvkicnchxccozrp
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The end of the time range to query. Specify the time in the 2006-01-02 format. By default, the time range to query is not restricted.
	//
	// example:
	//
	// 2022-02-01
	MaxDate *string `json:"MaxDate,omitempty" xml:"MaxDate,omitempty"`
	// The beginning of the time range to query. Specify the time in the 2006-01-02 format. By default, the time range to query is not restricted.
	//
	// example:
	//
	// 2022-01-02
	MinDate *string `json:"MinDate,omitempty" xml:"MinDate,omitempty"`
	// The page number. Pages start from page 1. This parameter is optional if you want to return the push status of all data files.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. This parameter is optional if you want to return the distribution status of all data files.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDataDistResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeDataDistResultShrinkRequest) GetDataNames() *string {
	return s.DataNames
}

func (s *DescribeDataDistResultShrinkRequest) GetDataVersions() *string {
	return s.DataVersions
}

func (s *DescribeDataDistResultShrinkRequest) GetEnsRegionIdsShrink() *string {
	return s.EnsRegionIdsShrink
}

func (s *DescribeDataDistResultShrinkRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeDataDistResultShrinkRequest) GetMaxDate() *string {
	return s.MaxDate
}

func (s *DescribeDataDistResultShrinkRequest) GetMinDate() *string {
	return s.MinDate
}

func (s *DescribeDataDistResultShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataDistResultShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataDistResultShrinkRequest) SetAppId(v string) *DescribeDataDistResultShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) SetDataNames(v string) *DescribeDataDistResultShrinkRequest {
	s.DataNames = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) SetDataVersions(v string) *DescribeDataDistResultShrinkRequest {
	s.DataVersions = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) SetEnsRegionIdsShrink(v string) *DescribeDataDistResultShrinkRequest {
	s.EnsRegionIdsShrink = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) SetInstanceIds(v string) *DescribeDataDistResultShrinkRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) SetMaxDate(v string) *DescribeDataDistResultShrinkRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) SetMinDate(v string) *DescribeDataDistResultShrinkRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) SetPageNumber(v int32) *DescribeDataDistResultShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) SetPageSize(v int32) *DescribeDataDistResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataDistResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
