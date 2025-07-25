// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorConfigId(v string) *AddGtmMonitorResponseBody
	GetMonitorConfigId() *string
	SetRequestId(v string) *AddGtmMonitorResponseBody
	GetRequestId() *string
}

type AddGtmMonitorResponseBody struct {
	// The ID of the health check configuration.
	//
	// example:
	//
	// 1234abc
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGtmMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGtmMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *AddGtmMonitorResponseBody) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *AddGtmMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGtmMonitorResponseBody) SetMonitorConfigId(v string) *AddGtmMonitorResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *AddGtmMonitorResponseBody) SetRequestId(v string) *AddGtmMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGtmMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
