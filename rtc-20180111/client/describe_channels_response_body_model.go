// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *DescribeChannelsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeChannelsResponseBody
	GetPageSize() *int32
	SetRecords(v []*string) *DescribeChannelsResponseBody
	GetRecords() []*string
	SetRequestId(v string) *DescribeChannelsResponseBody
	GetRequestId() *string
	SetTotalCnt(v int32) *DescribeChannelsResponseBody
	GetTotalCnt() *int32
}

type DescribeChannelsResponseBody struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*string `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCnt *int32 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeChannelsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeChannelsResponseBody) GetRecords() []*string {
	return s.Records
}

func (s *DescribeChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelsResponseBody) GetTotalCnt() *int32 {
	return s.TotalCnt
}

func (s *DescribeChannelsResponseBody) SetPageNo(v int32) *DescribeChannelsResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeChannelsResponseBody) SetPageSize(v int32) *DescribeChannelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeChannelsResponseBody) SetRecords(v []*string) *DescribeChannelsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeChannelsResponseBody) SetRequestId(v string) *DescribeChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelsResponseBody) SetTotalCnt(v int32) *DescribeChannelsResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeChannelsResponseBody) Validate() error {
	return dara.Validate(s)
}
