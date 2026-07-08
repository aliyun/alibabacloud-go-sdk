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
	SetStatus(v bool) *DescribeSlsAuthStatusResponseBody
	GetStatus() *bool
}

type DescribeSlsAuthStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3589D3A3-4A04-51CB-AA89-353ED20ACB10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The authorization status. Valid values:
	//
	// - **true**: indicates that authorization is granted.
	//
	// - **false**: indicates that authorization is not granted.
	//
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeSlsAuthStatusResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DescribeSlsAuthStatusResponseBody) SetRequestId(v string) *DescribeSlsAuthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsAuthStatusResponseBody) SetStatus(v bool) *DescribeSlsAuthStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSlsAuthStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
