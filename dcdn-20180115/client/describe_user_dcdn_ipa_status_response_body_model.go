// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDcdnIpaStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *DescribeUserDcdnIpaStatusResponseBody
	GetEnabled() *bool
	SetInDebt(v bool) *DescribeUserDcdnIpaStatusResponseBody
	GetInDebt() *bool
	SetInDebtOverdue(v bool) *DescribeUserDcdnIpaStatusResponseBody
	GetInDebtOverdue() *bool
	SetOnService(v bool) *DescribeUserDcdnIpaStatusResponseBody
	GetOnService() *bool
	SetRequestId(v string) *DescribeUserDcdnIpaStatusResponseBody
	GetRequestId() *string
}

type DescribeUserDcdnIpaStatusResponseBody struct {
	// Indicates whether the IPA service is activated.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether you have overdue payments.
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
	// Indicates whether the IPA service is available. The IPA service is available when no payment is overdue.
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

func (s DescribeUserDcdnIpaStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDcdnIpaStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnIpaStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeUserDcdnIpaStatusResponseBody) GetInDebt() *bool {
	return s.InDebt
}

func (s *DescribeUserDcdnIpaStatusResponseBody) GetInDebtOverdue() *bool {
	return s.InDebtOverdue
}

func (s *DescribeUserDcdnIpaStatusResponseBody) GetOnService() *bool {
	return s.OnService
}

func (s *DescribeUserDcdnIpaStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetEnabled(v bool) *DescribeUserDcdnIpaStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetInDebt(v bool) *DescribeUserDcdnIpaStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserDcdnIpaStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetOnService(v bool) *DescribeUserDcdnIpaStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) SetRequestId(v string) *DescribeUserDcdnIpaStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserDcdnIpaStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
