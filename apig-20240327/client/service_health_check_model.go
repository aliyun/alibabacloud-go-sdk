// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceHealthCheck interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ServiceHealthCheck
	GetEnable() *bool
	SetHealthyThreshold(v int32) *ServiceHealthCheck
	GetHealthyThreshold() *int32
	SetHttpHost(v string) *ServiceHealthCheck
	GetHttpHost() *string
	SetHttpPath(v string) *ServiceHealthCheck
	GetHttpPath() *string
	SetInterval(v int32) *ServiceHealthCheck
	GetInterval() *int32
	SetProtocol(v string) *ServiceHealthCheck
	GetProtocol() *string
	SetTimeout(v int32) *ServiceHealthCheck
	GetTimeout() *int32
	SetUnhealthyThreshold(v int32) *ServiceHealthCheck
	GetUnhealthyThreshold() *int32
}

type ServiceHealthCheck struct {
	// example:
	//
	// true
	Enable           *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	HealthyThreshold *int32  `json:"healthyThreshold,omitempty" xml:"healthyThreshold,omitempty"`
	HttpHost         *string `json:"httpHost,omitempty" xml:"httpHost,omitempty"`
	HttpPath         *string `json:"httpPath,omitempty" xml:"httpPath,omitempty"`
	Interval         *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// TCP
	Protocol           *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	Timeout            *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	UnhealthyThreshold *int32  `json:"unhealthyThreshold,omitempty" xml:"unhealthyThreshold,omitempty"`
}

func (s ServiceHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s ServiceHealthCheck) GoString() string {
	return s.String()
}

func (s *ServiceHealthCheck) GetEnable() *bool {
	return s.Enable
}

func (s *ServiceHealthCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *ServiceHealthCheck) GetHttpHost() *string {
	return s.HttpHost
}

func (s *ServiceHealthCheck) GetHttpPath() *string {
	return s.HttpPath
}

func (s *ServiceHealthCheck) GetInterval() *int32 {
	return s.Interval
}

func (s *ServiceHealthCheck) GetProtocol() *string {
	return s.Protocol
}

func (s *ServiceHealthCheck) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ServiceHealthCheck) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *ServiceHealthCheck) SetEnable(v bool) *ServiceHealthCheck {
	s.Enable = &v
	return s
}

func (s *ServiceHealthCheck) SetHealthyThreshold(v int32) *ServiceHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *ServiceHealthCheck) SetHttpHost(v string) *ServiceHealthCheck {
	s.HttpHost = &v
	return s
}

func (s *ServiceHealthCheck) SetHttpPath(v string) *ServiceHealthCheck {
	s.HttpPath = &v
	return s
}

func (s *ServiceHealthCheck) SetInterval(v int32) *ServiceHealthCheck {
	s.Interval = &v
	return s
}

func (s *ServiceHealthCheck) SetProtocol(v string) *ServiceHealthCheck {
	s.Protocol = &v
	return s
}

func (s *ServiceHealthCheck) SetTimeout(v int32) *ServiceHealthCheck {
	s.Timeout = &v
	return s
}

func (s *ServiceHealthCheck) SetUnhealthyThreshold(v int32) *ServiceHealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *ServiceHealthCheck) Validate() error {
	return dara.Validate(s)
}
