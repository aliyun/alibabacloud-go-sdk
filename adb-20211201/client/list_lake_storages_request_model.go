// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLakeStoragesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListLakeStoragesRequest
	GetDBClusterId() *string
	SetFilter(v string) *ListLakeStoragesRequest
	GetFilter() *string
	SetPageNumber(v int32) *ListLakeStoragesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLakeStoragesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListLakeStoragesRequest
	GetRegionId() *string
}

type ListLakeStoragesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The filter parameters that you want to use to query lake storages. Specify multiple parameters in an AND relationship. For example, if you want to query lake storage whose names are in the range of i-a123, or i-b123, and in the Stopped state, set this parameter to \\&Filter. 1.Name=InstanceName\\&Filter. 1.Value.1=i-a123\\&Filter.1.Value.2=i-b123\\&Filter.2.Name=Status\\&Filter. 2.Value=Stopped.
	//
	// example:
	//
	// -
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLakeStoragesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLakeStoragesRequest) GoString() string {
	return s.String()
}

func (s *ListLakeStoragesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListLakeStoragesRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListLakeStoragesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLakeStoragesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLakeStoragesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLakeStoragesRequest) SetDBClusterId(v string) *ListLakeStoragesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListLakeStoragesRequest) SetFilter(v string) *ListLakeStoragesRequest {
	s.Filter = &v
	return s
}

func (s *ListLakeStoragesRequest) SetPageNumber(v int32) *ListLakeStoragesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLakeStoragesRequest) SetPageSize(v int32) *ListLakeStoragesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLakeStoragesRequest) SetRegionId(v string) *ListLakeStoragesRequest {
	s.RegionId = &v
	return s
}

func (s *ListLakeStoragesRequest) Validate() error {
	return dara.Validate(s)
}
