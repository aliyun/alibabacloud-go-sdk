// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableCrossRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*string) *DescribeAvailableCrossRegionsResponseBody
	GetRegions() []*string
	SetRequestId(v string) *DescribeAvailableCrossRegionsResponseBody
	GetRequestId() *string
}

type DescribeAvailableCrossRegionsResponseBody struct {
	// example:
	//
	// [\\"cn-beijing\\", \\"cn-shanghai\\", \\"cn-qingdao\\", \\"cn-shenzhen\\", \\"cn-hongkong\\"]
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// D685D479-B36E-52B9-98FF-8402EA01F***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableCrossRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionsResponseBody) GetRegions() []*string {
	return s.Regions
}

func (s *DescribeAvailableCrossRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableCrossRegionsResponseBody) SetRegions(v []*string) *DescribeAvailableCrossRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeAvailableCrossRegionsResponseBody) SetRequestId(v string) *DescribeAvailableCrossRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableCrossRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
