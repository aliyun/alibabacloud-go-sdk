// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPoolId(v string) *AddDnsGtmAddressPoolResponseBody
	GetAddrPoolId() *string
	SetMonitorConfigId(v string) *AddDnsGtmAddressPoolResponseBody
	GetMonitorConfigId() *string
	SetRequestId(v string) *AddDnsGtmAddressPoolResponseBody
	GetRequestId() *string
}

type AddDnsGtmAddressPoolResponseBody struct {
	// The ID of the address pool.
	//
	// example:
	//
	// testpool1
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The ID of the health check configuration.
	//
	// example:
	//
	// test1
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDnsGtmAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolResponseBody) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *AddDnsGtmAddressPoolResponseBody) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *AddDnsGtmAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDnsGtmAddressPoolResponseBody) SetAddrPoolId(v string) *AddDnsGtmAddressPoolResponseBody {
	s.AddrPoolId = &v
	return s
}

func (s *AddDnsGtmAddressPoolResponseBody) SetMonitorConfigId(v string) *AddDnsGtmAddressPoolResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *AddDnsGtmAddressPoolResponseBody) SetRequestId(v string) *AddDnsGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnsGtmAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
