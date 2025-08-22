// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserErStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *DescribeUserErStatusResponseBody
	GetEnabled() *bool
	SetInDebt(v bool) *DescribeUserErStatusResponseBody
	GetInDebt() *bool
	SetInDebtOverdue(v bool) *DescribeUserErStatusResponseBody
	GetInDebtOverdue() *bool
	SetOnService(v bool) *DescribeUserErStatusResponseBody
	GetOnService() *bool
	SetRequestId(v string) *DescribeUserErStatusResponseBody
	GetRequestId() *string
}

type DescribeUserErStatusResponseBody struct {
	// Indicates whether ER is activated.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether ER has an overdue payment.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	InDebt *bool `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	// Indicates whether an overdue payment of ER has passed the grace period.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	InDebtOverdue *bool `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	// Indicates whether ER is available.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	OnService *bool `json:"OnService,omitempty" xml:"OnService,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F51E9C3-728F-4E35-952D-0ED87A06A8A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserErStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserErStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserErStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeUserErStatusResponseBody) GetInDebt() *bool {
	return s.InDebt
}

func (s *DescribeUserErStatusResponseBody) GetInDebtOverdue() *bool {
	return s.InDebtOverdue
}

func (s *DescribeUserErStatusResponseBody) GetOnService() *bool {
	return s.OnService
}

func (s *DescribeUserErStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserErStatusResponseBody) SetEnabled(v bool) *DescribeUserErStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) SetInDebt(v bool) *DescribeUserErStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserErStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) SetOnService(v bool) *DescribeUserErStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) SetRequestId(v string) *DescribeUserErStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserErStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
