// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniSupportRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUniSupportRegionResponseBody
	GetRequestId() *string
	SetUniSupportRegion(v []*string) *DescribeUniSupportRegionResponseBody
	GetUniSupportRegion() []*string
}

type DescribeUniSupportRegionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D0760****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the region that is supported by anti-ransomware for databases.
	UniSupportRegion []*string `json:"UniSupportRegion,omitempty" xml:"UniSupportRegion,omitempty" type:"Repeated"`
}

func (s DescribeUniSupportRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniSupportRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUniSupportRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUniSupportRegionResponseBody) GetUniSupportRegion() []*string {
	return s.UniSupportRegion
}

func (s *DescribeUniSupportRegionResponseBody) SetRequestId(v string) *DescribeUniSupportRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUniSupportRegionResponseBody) SetUniSupportRegion(v []*string) *DescribeUniSupportRegionResponseBody {
	s.UniSupportRegion = v
	return s
}

func (s *DescribeUniSupportRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
