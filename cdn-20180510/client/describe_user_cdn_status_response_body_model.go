// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserCdnStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *DescribeUserCdnStatusResponseBody
	GetEnabled() *bool
	SetInDebt(v bool) *DescribeUserCdnStatusResponseBody
	GetInDebt() *bool
	SetInDebtOverdue(v bool) *DescribeUserCdnStatusResponseBody
	GetInDebtOverdue() *bool
	SetOnService(v bool) *DescribeUserCdnStatusResponseBody
	GetOnService() *bool
	SetRequestId(v string) *DescribeUserCdnStatusResponseBody
	GetRequestId() *string
}

type DescribeUserCdnStatusResponseBody struct {
	// Indicates whetherAlibaba Cloud CDN is activated.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether your account has overdue payments.
	//
	// 	- true
	//
	// 	- false
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
	// The request ID.
	//
	// example:
	//
	// 39B1DC7F-9D25-5D54-8F02-6EE26A7F48CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserCdnStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserCdnStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserCdnStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeUserCdnStatusResponseBody) GetInDebt() *bool {
	return s.InDebt
}

func (s *DescribeUserCdnStatusResponseBody) GetInDebtOverdue() *bool {
	return s.InDebtOverdue
}

func (s *DescribeUserCdnStatusResponseBody) GetOnService() *bool {
	return s.OnService
}

func (s *DescribeUserCdnStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserCdnStatusResponseBody) SetEnabled(v bool) *DescribeUserCdnStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserCdnStatusResponseBody) SetInDebt(v bool) *DescribeUserCdnStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserCdnStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserCdnStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserCdnStatusResponseBody) SetOnService(v bool) *DescribeUserCdnStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserCdnStatusResponseBody) SetRequestId(v string) *DescribeUserCdnStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserCdnStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
