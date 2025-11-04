// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListSourcesResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListSourcesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSourcesResponseBody
	GetRequestId() *string
	SetSourceList(v []*ChannelAssemblySource) *ListSourcesResponseBody
	GetSourceList() []*ChannelAssemblySource
	SetTotalCount(v int32) *ListSourcesResponseBody
	GetTotalCount() *int32
}

type ListSourcesResponseBody struct {
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sources.
	SourceList []*ChannelAssemblySource `json:"SourceList,omitempty" xml:"SourceList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSourcesResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSourcesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSourcesResponseBody) GetSourceList() []*ChannelAssemblySource {
	return s.SourceList
}

func (s *ListSourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSourcesResponseBody) SetPageNo(v int32) *ListSourcesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListSourcesResponseBody) SetPageSize(v int32) *ListSourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSourcesResponseBody) SetRequestId(v string) *ListSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSourcesResponseBody) SetSourceList(v []*ChannelAssemblySource) *ListSourcesResponseBody {
	s.SourceList = v
	return s
}

func (s *ListSourcesResponseBody) SetTotalCount(v int32) *ListSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSourcesResponseBody) Validate() error {
	if s.SourceList != nil {
		for _, item := range s.SourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
