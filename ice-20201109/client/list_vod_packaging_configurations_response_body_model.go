// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackagingConfigurations(v []*VodPackagingConfiguration) *ListVodPackagingConfigurationsResponseBody
	GetPackagingConfigurations() []*VodPackagingConfiguration
	SetPageNo(v int64) *ListVodPackagingConfigurationsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListVodPackagingConfigurationsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListVodPackagingConfigurationsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListVodPackagingConfigurationsResponseBody
	GetSortBy() *string
	SetTotalCount(v int32) *ListVodPackagingConfigurationsResponseBody
	GetTotalCount() *int32
}

type ListVodPackagingConfigurationsResponseBody struct {
	// The packaging configurations.
	PackagingConfigurations []*VodPackagingConfiguration `json:"PackagingConfigurations,omitempty" xml:"PackagingConfigurations,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order of the packaging configurations based on the time when they were created. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVodPackagingConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodPackagingConfigurationsResponseBody) GetPackagingConfigurations() []*VodPackagingConfiguration {
	return s.PackagingConfigurations
}

func (s *ListVodPackagingConfigurationsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListVodPackagingConfigurationsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVodPackagingConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodPackagingConfigurationsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListVodPackagingConfigurationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVodPackagingConfigurationsResponseBody) SetPackagingConfigurations(v []*VodPackagingConfiguration) *ListVodPackagingConfigurationsResponseBody {
	s.PackagingConfigurations = v
	return s
}

func (s *ListVodPackagingConfigurationsResponseBody) SetPageNo(v int64) *ListVodPackagingConfigurationsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListVodPackagingConfigurationsResponseBody) SetPageSize(v int64) *ListVodPackagingConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVodPackagingConfigurationsResponseBody) SetRequestId(v string) *ListVodPackagingConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodPackagingConfigurationsResponseBody) SetSortBy(v string) *ListVodPackagingConfigurationsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListVodPackagingConfigurationsResponseBody) SetTotalCount(v int32) *ListVodPackagingConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVodPackagingConfigurationsResponseBody) Validate() error {
	if s.PackagingConfigurations != nil {
		for _, item := range s.PackagingConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
