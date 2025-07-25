// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorConfigId(v string) *AddDnsGtmMonitorResponseBody
	GetMonitorConfigId() *string
	SetRequestId(v string) *AddDnsGtmMonitorResponseBody
	GetRequestId() *string
}

type AddDnsGtmMonitorResponseBody struct {
	// The ID of the health check configuration.
	//
	// example:
	//
	// MonitorConfigId1
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDnsGtmMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnsGtmMonitorResponseBody) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *AddDnsGtmMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDnsGtmMonitorResponseBody) SetMonitorConfigId(v string) *AddDnsGtmMonitorResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *AddDnsGtmMonitorResponseBody) SetRequestId(v string) *AddDnsGtmMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnsGtmMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
