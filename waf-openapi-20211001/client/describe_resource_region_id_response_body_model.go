// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceRegionIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourceRegionIdResponseBody
	GetRequestId() *string
	SetResourceRegionIds(v []*string) *DescribeResourceRegionIdResponseBody
	GetResourceRegionIds() []*string
}

type DescribeResourceRegionIdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F5905D3F-F674-5177-9E48-466DD3B8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The region IDs of the resources that are added to Web Application Firewall (WAF) by using the SDK integration mode.
	ResourceRegionIds []*string `json:"ResourceRegionIds,omitempty" xml:"ResourceRegionIds,omitempty" type:"Repeated"`
}

func (s DescribeResourceRegionIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceRegionIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceRegionIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceRegionIdResponseBody) GetResourceRegionIds() []*string {
	return s.ResourceRegionIds
}

func (s *DescribeResourceRegionIdResponseBody) SetRequestId(v string) *DescribeResourceRegionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceRegionIdResponseBody) SetResourceRegionIds(v []*string) *DescribeResourceRegionIdResponseBody {
	s.ResourceRegionIds = v
	return s
}

func (s *DescribeResourceRegionIdResponseBody) Validate() error {
	return dara.Validate(s)
}
