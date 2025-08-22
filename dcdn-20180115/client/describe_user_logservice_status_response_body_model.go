// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogserviceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *DescribeUserLogserviceStatusResponseBody
	GetEnabled() *bool
	SetInDebt(v bool) *DescribeUserLogserviceStatusResponseBody
	GetInDebt() *bool
	SetInDebtOverdue(v bool) *DescribeUserLogserviceStatusResponseBody
	GetInDebtOverdue() *bool
	SetOnService(v bool) *DescribeUserLogserviceStatusResponseBody
	GetOnService() *bool
	SetRequestId(v string) *DescribeUserLogserviceStatusResponseBody
	GetRequestId() *string
}

type DescribeUserLogserviceStatusResponseBody struct {
	// Indicates whether Log Service is activated.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether your Log Service has overdue payments.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	InDebt *bool `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	// Indicates whether an overdue payment of your Log Service has passed the grace period.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	InDebtOverdue *bool `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	// Indicates whether Log Service is available.
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

func (s DescribeUserLogserviceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogserviceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserLogserviceStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeUserLogserviceStatusResponseBody) GetInDebt() *bool {
	return s.InDebt
}

func (s *DescribeUserLogserviceStatusResponseBody) GetInDebtOverdue() *bool {
	return s.InDebtOverdue
}

func (s *DescribeUserLogserviceStatusResponseBody) GetOnService() *bool {
	return s.OnService
}

func (s *DescribeUserLogserviceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserLogserviceStatusResponseBody) SetEnabled(v bool) *DescribeUserLogserviceStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) SetInDebt(v bool) *DescribeUserLogserviceStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserLogserviceStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) SetOnService(v bool) *DescribeUserLogserviceStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) SetRequestId(v string) *DescribeUserLogserviceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
