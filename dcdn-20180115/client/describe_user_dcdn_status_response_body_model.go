// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDcdnStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *DescribeUserDcdnStatusResponseBody
	GetEnabled() *bool
	SetInDebt(v bool) *DescribeUserDcdnStatusResponseBody
	GetInDebt() *bool
	SetInDebtOverdue(v bool) *DescribeUserDcdnStatusResponseBody
	GetInDebtOverdue() *bool
	SetOnService(v bool) *DescribeUserDcdnStatusResponseBody
	GetOnService() *bool
	SetRequestId(v string) *DescribeUserDcdnStatusResponseBody
	GetRequestId() *string
}

type DescribeUserDcdnStatusResponseBody struct {
	// Indicates whether the DCDN service is activated.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether your account has overdue payments.
	//
	// example:
	//
	// false
	InDebt *bool `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	// Indicates whether the grace period for your overdue payments expired.
	//
	// example:
	//
	// false
	InDebtOverdue *bool `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	// Indicates whether the service is available.
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

func (s DescribeUserDcdnStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDcdnStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeUserDcdnStatusResponseBody) GetInDebt() *bool {
	return s.InDebt
}

func (s *DescribeUserDcdnStatusResponseBody) GetInDebtOverdue() *bool {
	return s.InDebtOverdue
}

func (s *DescribeUserDcdnStatusResponseBody) GetOnService() *bool {
	return s.OnService
}

func (s *DescribeUserDcdnStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserDcdnStatusResponseBody) SetEnabled(v bool) *DescribeUserDcdnStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) SetInDebt(v bool) *DescribeUserDcdnStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserDcdnStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) SetOnService(v bool) *DescribeUserDcdnStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) SetRequestId(v string) *DescribeUserDcdnStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserDcdnStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
