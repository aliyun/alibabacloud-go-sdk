// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExceptionCount(v int64) *DescribeNatFirewallQuotaResponseBody
	GetExceptionCount() *int64
	SetRequestId(v string) *DescribeNatFirewallQuotaResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeNatFirewallQuotaResponseBody
	GetTotalCount() *int64
	SetUnprotectedCount(v int64) *DescribeNatFirewallQuotaResponseBody
	GetUnprotectedCount() *int64
	SetUsedCount(v int64) *DescribeNatFirewallQuotaResponseBody
	GetUsedCount() *int64
}

type DescribeNatFirewallQuotaResponseBody struct {
	// example:
	//
	// 1
	ExceptionCount *int64 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	// example:
	//
	// F98BAA59-5863-5B61-8FD4-C5E96813****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 20
	UnprotectedCount *int64 `json:"UnprotectedCount,omitempty" xml:"UnprotectedCount,omitempty"`
	// example:
	//
	// 10
	UsedCount *int64 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeNatFirewallQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallQuotaResponseBody) GetExceptionCount() *int64 {
	return s.ExceptionCount
}

func (s *DescribeNatFirewallQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallQuotaResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeNatFirewallQuotaResponseBody) GetUnprotectedCount() *int64 {
	return s.UnprotectedCount
}

func (s *DescribeNatFirewallQuotaResponseBody) GetUsedCount() *int64 {
	return s.UsedCount
}

func (s *DescribeNatFirewallQuotaResponseBody) SetExceptionCount(v int64) *DescribeNatFirewallQuotaResponseBody {
	s.ExceptionCount = &v
	return s
}

func (s *DescribeNatFirewallQuotaResponseBody) SetRequestId(v string) *DescribeNatFirewallQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallQuotaResponseBody) SetTotalCount(v int64) *DescribeNatFirewallQuotaResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNatFirewallQuotaResponseBody) SetUnprotectedCount(v int64) *DescribeNatFirewallQuotaResponseBody {
	s.UnprotectedCount = &v
	return s
}

func (s *DescribeNatFirewallQuotaResponseBody) SetUsedCount(v int64) *DescribeNatFirewallQuotaResponseBody {
	s.UsedCount = &v
	return s
}

func (s *DescribeNatFirewallQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
