// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUdpCheck interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckConnectPort(v int32) *UdpCheck
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *UdpCheck
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckInterval(v int32) *UdpCheck
	GetHealthCheckInterval() *int32
	SetHealthyThreshold(v int32) *UdpCheck
	GetHealthyThreshold() *int32
	SetUnhealthyThreshold(v int32) *UdpCheck
	GetUnhealthyThreshold() *int32
}

type UdpCheck struct {
	HealthCheckConnectPort    *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckInterval       *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthyThreshold          *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	UnhealthyThreshold        *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UdpCheck) String() string {
	return dara.Prettify(s)
}

func (s UdpCheck) GoString() string {
	return s.String()
}

func (s *UdpCheck) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *UdpCheck) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *UdpCheck) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *UdpCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UdpCheck) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *UdpCheck) SetHealthCheckConnectPort(v int32) *UdpCheck {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *UdpCheck) SetHealthCheckConnectTimeout(v int32) *UdpCheck {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *UdpCheck) SetHealthCheckInterval(v int32) *UdpCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *UdpCheck) SetHealthyThreshold(v int32) *UdpCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *UdpCheck) SetUnhealthyThreshold(v int32) *UdpCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UdpCheck) Validate() error {
	return dara.Validate(s)
}
