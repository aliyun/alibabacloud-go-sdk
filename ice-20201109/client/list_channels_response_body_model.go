// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelList(v []*ChannelAssemblyChannel) *ListChannelsResponseBody
	GetChannelList() []*ChannelAssemblyChannel
	SetPageNo(v int32) *ListChannelsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListChannelsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListChannelsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListChannelsResponseBody
	GetTotalCount() *int32
}

type ListChannelsResponseBody struct {
	// The channels.
	ChannelList []*ChannelAssemblyChannel `json:"ChannelList,omitempty" xml:"ChannelList,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
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
	// The total number of channels returned.
	//
	// example:
	//
	// 180
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChannelsResponseBody) GetChannelList() []*ChannelAssemblyChannel {
	return s.ChannelList
}

func (s *ListChannelsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChannelsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChannelsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListChannelsResponseBody) SetChannelList(v []*ChannelAssemblyChannel) *ListChannelsResponseBody {
	s.ChannelList = v
	return s
}

func (s *ListChannelsResponseBody) SetPageNo(v int32) *ListChannelsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChannelsResponseBody) SetPageSize(v int32) *ListChannelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChannelsResponseBody) SetRequestId(v string) *ListChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChannelsResponseBody) SetTotalCount(v int32) *ListChannelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChannelsResponseBody) Validate() error {
	if s.ChannelList != nil {
		for _, item := range s.ChannelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
