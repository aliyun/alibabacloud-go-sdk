// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPoolId(v string) *AddGtmAddressPoolResponseBody
	GetAddrPoolId() *string
	SetMonitorConfigId(v string) *AddGtmAddressPoolResponseBody
	GetMonitorConfigId() *string
	SetRequestId(v string) *AddGtmAddressPoolResponseBody
	GetRequestId() *string
}

type AddGtmAddressPoolResponseBody struct {
	// The ID of the address pool created.
	//
	// example:
	//
	// hraf3x
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The ID of the health check configuration.
	//
	// example:
	//
	// hraf14
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGtmAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolResponseBody) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *AddGtmAddressPoolResponseBody) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *AddGtmAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGtmAddressPoolResponseBody) SetAddrPoolId(v string) *AddGtmAddressPoolResponseBody {
	s.AddrPoolId = &v
	return s
}

func (s *AddGtmAddressPoolResponseBody) SetMonitorConfigId(v string) *AddGtmAddressPoolResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *AddGtmAddressPoolResponseBody) SetRequestId(v string) *AddGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGtmAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
