// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceRegionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionNoList(v []*string) *DescribeAccessInstanceRegionListResponseBody
	GetRegionNoList() []*string
	SetRequestId(v string) *DescribeAccessInstanceRegionListResponseBody
	GetRequestId() *string
}

type DescribeAccessInstanceRegionListResponseBody struct {
	RegionNoList []*string `json:"RegionNoList,omitempty" xml:"RegionNoList,omitempty" type:"Repeated"`
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccessInstanceRegionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceRegionListResponseBody) GetRegionNoList() []*string {
	return s.RegionNoList
}

func (s *DescribeAccessInstanceRegionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessInstanceRegionListResponseBody) SetRegionNoList(v []*string) *DescribeAccessInstanceRegionListResponseBody {
	s.RegionNoList = v
	return s
}

func (s *DescribeAccessInstanceRegionListResponseBody) SetRequestId(v string) *DescribeAccessInstanceRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessInstanceRegionListResponseBody) Validate() error {
	return dara.Validate(s)
}
