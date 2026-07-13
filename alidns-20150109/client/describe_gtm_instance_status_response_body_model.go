// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody
	GetAddrNotAvailableNum() *int32
	SetAddrPoolNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody
	GetAddrPoolNotAvailableNum() *int32
	SetRequestId(v string) *DescribeGtmInstanceStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeGtmInstanceStatusResponseBody
	GetStatus() *string
	SetStatusReason(v string) *DescribeGtmInstanceStatusResponseBody
	GetStatusReason() *string
	SetStrategyNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody
	GetStrategyNotAvailableNum() *int32
	SetSwitchToFailoverStrategyNum(v int32) *DescribeGtmInstanceStatusResponseBody
	GetSwitchToFailoverStrategyNum() *int32
}

type DescribeGtmInstanceStatusResponseBody struct {
	// The number of unavailable addresses.
	//
	// example:
	//
	// 10
	AddrNotAvailableNum *int32 `json:"AddrNotAvailableNum,omitempty" xml:"AddrNotAvailableNum,omitempty"`
	// The number of unavailable address pools.
	//
	// example:
	//
	// 10
	AddrPoolNotAvailableNum *int32 `json:"AddrPoolNotAvailableNum,omitempty" xml:"AddrPoolNotAvailableNum,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 389DFFA3-77A5-4A9E-BF3D-147C6F98A5BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the instance. Valid values:
	//
	// - ALLOW: Operations are allowed.
	//
	// - DENY: Operations are denied.
	//
	// example:
	//
	// ALLOW
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of reasons for the instance status. Valid values:
	//
	// - INSTANCE_OPERATE_BLACK_LIST: The instance is in a blacklist.
	//
	// - BETA_INSTANCE: The instance is in public preview.
	//
	// example:
	//
	// ["BETA_INSTANCE"]
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The number of access policies for which the active address pool is unavailable.
	//
	// example:
	//
	// 10
	StrategyNotAvailableNum *int32 `json:"StrategyNotAvailableNum,omitempty" xml:"StrategyNotAvailableNum,omitempty"`
	// The number of access policies that are switched to the failover address pool.
	//
	// example:
	//
	// 10
	SwitchToFailoverStrategyNum *int32 `json:"SwitchToFailoverStrategyNum,omitempty" xml:"SwitchToFailoverStrategyNum,omitempty"`
}

func (s DescribeGtmInstanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceStatusResponseBody) GetAddrNotAvailableNum() *int32 {
	return s.AddrNotAvailableNum
}

func (s *DescribeGtmInstanceStatusResponseBody) GetAddrPoolNotAvailableNum() *int32 {
	return s.AddrPoolNotAvailableNum
}

func (s *DescribeGtmInstanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmInstanceStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGtmInstanceStatusResponseBody) GetStatusReason() *string {
	return s.StatusReason
}

func (s *DescribeGtmInstanceStatusResponseBody) GetStrategyNotAvailableNum() *int32 {
	return s.StrategyNotAvailableNum
}

func (s *DescribeGtmInstanceStatusResponseBody) GetSwitchToFailoverStrategyNum() *int32 {
	return s.SwitchToFailoverStrategyNum
}

func (s *DescribeGtmInstanceStatusResponseBody) SetAddrNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody {
	s.AddrNotAvailableNum = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetAddrPoolNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody {
	s.AddrPoolNotAvailableNum = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetRequestId(v string) *DescribeGtmInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetStatus(v string) *DescribeGtmInstanceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetStatusReason(v string) *DescribeGtmInstanceStatusResponseBody {
	s.StatusReason = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetStrategyNotAvailableNum(v int32) *DescribeGtmInstanceStatusResponseBody {
	s.StrategyNotAvailableNum = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) SetSwitchToFailoverStrategyNum(v int32) *DescribeGtmInstanceStatusResponseBody {
	s.SwitchToFailoverStrategyNum = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
