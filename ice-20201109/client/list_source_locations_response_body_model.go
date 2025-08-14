// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceLocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListSourceLocationsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListSourceLocationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSourceLocationsResponseBody
	GetRequestId() *string
	SetSourceLocationList(v []*ChannelAssemblySourceLocation) *ListSourceLocationsResponseBody
	GetSourceLocationList() []*ChannelAssemblySourceLocation
	SetTotalCount(v int32) *ListSourceLocationsResponseBody
	GetTotalCount() *int32
}

type ListSourceLocationsResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source locations.
	SourceLocationList []*ChannelAssemblySourceLocation `json:"SourceLocationList,omitempty" xml:"SourceLocationList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSourceLocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSourceLocationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSourceLocationsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSourceLocationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSourceLocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSourceLocationsResponseBody) GetSourceLocationList() []*ChannelAssemblySourceLocation {
	return s.SourceLocationList
}

func (s *ListSourceLocationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSourceLocationsResponseBody) SetPageNo(v int32) *ListSourceLocationsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListSourceLocationsResponseBody) SetPageSize(v int32) *ListSourceLocationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSourceLocationsResponseBody) SetRequestId(v string) *ListSourceLocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSourceLocationsResponseBody) SetSourceLocationList(v []*ChannelAssemblySourceLocation) *ListSourceLocationsResponseBody {
	s.SourceLocationList = v
	return s
}

func (s *ListSourceLocationsResponseBody) SetTotalCount(v int32) *ListSourceLocationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSourceLocationsResponseBody) Validate() error {
	return dara.Validate(s)
}
