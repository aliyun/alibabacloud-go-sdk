// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppVulScanCycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCycle(v string) *DescribeAppVulScanCycleResponseBody
	GetCycle() *string
	SetRequestId(v string) *DescribeAppVulScanCycleResponseBody
	GetRequestId() *string
}

type DescribeAppVulScanCycleResponseBody struct {
	// The scan cycle for application vulnerabilities.
	//
	// 	- 1week
	//
	// 	- 2weeks
	//
	// 	- 3days
	//
	// example:
	//
	// 1week
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 934E6D2A-0123-5A99-88BA-80DC27634E22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppVulScanCycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppVulScanCycleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppVulScanCycleResponseBody) GetCycle() *string {
	return s.Cycle
}

func (s *DescribeAppVulScanCycleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppVulScanCycleResponseBody) SetCycle(v string) *DescribeAppVulScanCycleResponseBody {
	s.Cycle = &v
	return s
}

func (s *DescribeAppVulScanCycleResponseBody) SetRequestId(v string) *DescribeAppVulScanCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppVulScanCycleResponseBody) Validate() error {
	return dara.Validate(s)
}
