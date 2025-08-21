// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataDistResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeDataDistResultRequest
	GetAppId() *string
	SetDataNames(v string) *DescribeDataDistResultRequest
	GetDataNames() *string
	SetDataVersions(v string) *DescribeDataDistResultRequest
	GetDataVersions() *string
	SetEnsRegionIds(v []*string) *DescribeDataDistResultRequest
	GetEnsRegionIds() []*string
	SetInstanceIds(v string) *DescribeDataDistResultRequest
	GetInstanceIds() *string
	SetMaxDate(v string) *DescribeDataDistResultRequest
	GetMaxDate() *string
	SetMinDate(v string) *DescribeDataDistResultRequest
	GetMinDate() *string
	SetPageNumber(v int32) *DescribeDataDistResultRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataDistResultRequest
	GetPageSize() *int32
}

type DescribeDataDistResultRequest struct {
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
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
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

func (s DescribeDataDistResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeDataDistResultRequest) GetDataNames() *string {
	return s.DataNames
}

func (s *DescribeDataDistResultRequest) GetDataVersions() *string {
	return s.DataVersions
}

func (s *DescribeDataDistResultRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeDataDistResultRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeDataDistResultRequest) GetMaxDate() *string {
	return s.MaxDate
}

func (s *DescribeDataDistResultRequest) GetMinDate() *string {
	return s.MinDate
}

func (s *DescribeDataDistResultRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataDistResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataDistResultRequest) SetAppId(v string) *DescribeDataDistResultRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetDataNames(v string) *DescribeDataDistResultRequest {
	s.DataNames = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetDataVersions(v string) *DescribeDataDistResultRequest {
	s.DataVersions = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetEnsRegionIds(v []*string) *DescribeDataDistResultRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeDataDistResultRequest) SetInstanceIds(v string) *DescribeDataDistResultRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetMaxDate(v string) *DescribeDataDistResultRequest {
	s.MaxDate = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetMinDate(v string) *DescribeDataDistResultRequest {
	s.MinDate = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetPageNumber(v int32) *DescribeDataDistResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataDistResultRequest) SetPageSize(v int32) *DescribeDataDistResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataDistResultRequest) Validate() error {
	return dara.Validate(s)
}
