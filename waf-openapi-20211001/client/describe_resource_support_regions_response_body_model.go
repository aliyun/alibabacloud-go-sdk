// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceSupportRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourceSupportRegionsResponseBody
	GetRequestId() *string
	SetSupportRegions(v []*string) *DescribeResourceSupportRegionsResponseBody
	GetSupportRegions() []*string
}

type DescribeResourceSupportRegionsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 58FD****-3D56-5DE8-91E0-96A26BABFFDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of region IDs of the CLB and ECS instances that are added to WAF in cloud native mode.
	SupportRegions []*string `json:"SupportRegions,omitempty" xml:"SupportRegions,omitempty" type:"Repeated"`
}

func (s DescribeResourceSupportRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceSupportRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceSupportRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceSupportRegionsResponseBody) GetSupportRegions() []*string {
	return s.SupportRegions
}

func (s *DescribeResourceSupportRegionsResponseBody) SetRequestId(v string) *DescribeResourceSupportRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceSupportRegionsResponseBody) SetSupportRegions(v []*string) *DescribeResourceSupportRegionsResponseBody {
	s.SupportRegions = v
	return s
}

func (s *DescribeResourceSupportRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
