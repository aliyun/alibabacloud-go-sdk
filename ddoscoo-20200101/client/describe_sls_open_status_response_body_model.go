// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsOpenStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSlsOpenStatusResponseBody
	GetRequestId() *string
	SetSlsOpenStatus(v bool) *DescribeSlsOpenStatusResponseBody
	GetSlsOpenStatus() *bool
}

type DescribeSlsOpenStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Log Service is activated. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SlsOpenStatus *bool `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty"`
}

func (s DescribeSlsOpenStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsOpenStatusResponseBody) GetSlsOpenStatus() *bool {
	return s.SlsOpenStatus
}

func (s *DescribeSlsOpenStatusResponseBody) SetRequestId(v string) *DescribeSlsOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsOpenStatusResponseBody) SetSlsOpenStatus(v bool) *DescribeSlsOpenStatusResponseBody {
	s.SlsOpenStatus = &v
	return s
}

func (s *DescribeSlsOpenStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
