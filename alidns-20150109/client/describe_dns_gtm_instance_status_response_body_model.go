// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody
	GetAddrAvailableNum() *int32
	SetAddrNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody
	GetAddrNotAvailableNum() *int32
	SetAddrPoolGroupNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody
	GetAddrPoolGroupNotAvailableNum() *int32
	SetRequestId(v string) *DescribeDnsGtmInstanceStatusResponseBody
	GetRequestId() *string
	SetStrategyNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody
	GetStrategyNotAvailableNum() *int32
	SetSwitchToFailoverStrategyNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody
	GetSwitchToFailoverStrategyNum() *int32
}

type DescribeDnsGtmInstanceStatusResponseBody struct {
	// The number of available addresses.
	//
	// example:
	//
	// 1
	AddrAvailableNum *int32 `json:"AddrAvailableNum,omitempty" xml:"AddrAvailableNum,omitempty"`
	// The number of unavailable addresses.
	//
	// example:
	//
	// 1
	AddrNotAvailableNum *int32 `json:"AddrNotAvailableNum,omitempty" xml:"AddrNotAvailableNum,omitempty"`
	// The number of unavailable address pool groups.
	//
	// example:
	//
	// 1
	AddrPoolGroupNotAvailableNum *int32 `json:"AddrPoolGroupNotAvailableNum,omitempty" xml:"AddrPoolGroupNotAvailableNum,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 389DFFA3-77A5-4A9E-BF3D-147C6F98A5BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of access policies that are unavailable in the current active address pool group.
	//
	// example:
	//
	// 1
	StrategyNotAvailableNum *int32 `json:"StrategyNotAvailableNum,omitempty" xml:"StrategyNotAvailableNum,omitempty"`
	// The number of access policies switched to the secondary address pool group.
	//
	// example:
	//
	// 1
	SwitchToFailoverStrategyNum *int32 `json:"SwitchToFailoverStrategyNum,omitempty" xml:"SwitchToFailoverStrategyNum,omitempty"`
}

func (s DescribeDnsGtmInstanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) GetAddrAvailableNum() *int32 {
	return s.AddrAvailableNum
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) GetAddrNotAvailableNum() *int32 {
	return s.AddrNotAvailableNum
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) GetAddrPoolGroupNotAvailableNum() *int32 {
	return s.AddrPoolGroupNotAvailableNum
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) GetStrategyNotAvailableNum() *int32 {
	return s.StrategyNotAvailableNum
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) GetSwitchToFailoverStrategyNum() *int32 {
	return s.SwitchToFailoverStrategyNum
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetAddrAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.AddrAvailableNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetAddrNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.AddrNotAvailableNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetAddrPoolGroupNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.AddrPoolGroupNotAvailableNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetStrategyNotAvailableNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.StrategyNotAvailableNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) SetSwitchToFailoverStrategyNum(v int32) *DescribeDnsGtmInstanceStatusResponseBody {
	s.SwitchToFailoverStrategyNum = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
