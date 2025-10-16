// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceRegionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceStatus(v string) *DescribeAccessInstanceRegionListRequest
	GetAccessInstanceStatus() *string
	SetAccessInstanceType(v string) *DescribeAccessInstanceRegionListRequest
	GetAccessInstanceType() *string
}

type DescribeAccessInstanceRegionListRequest struct {
	// example:
	//
	// ready
	AccessInstanceStatus *string `json:"AccessInstanceStatus,omitempty" xml:"AccessInstanceStatus,omitempty"`
	// example:
	//
	// AckClusterConnector
	AccessInstanceType *string `json:"AccessInstanceType,omitempty" xml:"AccessInstanceType,omitempty"`
}

func (s DescribeAccessInstanceRegionListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceRegionListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceRegionListRequest) GetAccessInstanceStatus() *string {
	return s.AccessInstanceStatus
}

func (s *DescribeAccessInstanceRegionListRequest) GetAccessInstanceType() *string {
	return s.AccessInstanceType
}

func (s *DescribeAccessInstanceRegionListRequest) SetAccessInstanceStatus(v string) *DescribeAccessInstanceRegionListRequest {
	s.AccessInstanceStatus = &v
	return s
}

func (s *DescribeAccessInstanceRegionListRequest) SetAccessInstanceType(v string) *DescribeAccessInstanceRegionListRequest {
	s.AccessInstanceType = &v
	return s
}

func (s *DescribeAccessInstanceRegionListRequest) Validate() error {
	return dara.Validate(s)
}
