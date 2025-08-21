// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsAuthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSlsAuthStatusResponseBody
	GetRequestId() *string
	SetSlsAuthStatus(v bool) *DescribeSlsAuthStatusResponseBody
	GetSlsAuthStatus() *bool
}

type DescribeSlsAuthStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Anti-DDoS Pro or Anti-DDoS Premium is authorized to access Log Service. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SlsAuthStatus *bool `json:"SlsAuthStatus,omitempty" xml:"SlsAuthStatus,omitempty"`
}

func (s DescribeSlsAuthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsAuthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsAuthStatusResponseBody) GetSlsAuthStatus() *bool {
	return s.SlsAuthStatus
}

func (s *DescribeSlsAuthStatusResponseBody) SetRequestId(v string) *DescribeSlsAuthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsAuthStatusResponseBody) SetSlsAuthStatus(v bool) *DescribeSlsAuthStatusResponseBody {
	s.SlsAuthStatus = &v
	return s
}

func (s *DescribeSlsAuthStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
