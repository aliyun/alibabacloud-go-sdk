// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserServiceStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeUserServiceStatusResponseBody
	GetStatus() *string
}

type DescribeUserServiceStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 99626905-678A-4E8A-984E-6AEB09993996
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current user\\"s service status:
	//
	// 	- **CLOSED**: Not activated
	//
	// 	- **OPENED**: Activated
	//
	// 	- **IN_DEBT**: Overdue payment
	//
	// 	- **IN_DEBT_OVER_DUE**: Payment overdue
	//
	// example:
	//
	// OPENED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserServiceStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeUserServiceStatusResponseBody) SetRequestId(v string) *DescribeUserServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserServiceStatusResponseBody) SetStatus(v string) *DescribeUserServiceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeUserServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
