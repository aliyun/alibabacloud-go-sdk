// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHealthCheckConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCheckInterval(v int32) *HealthCheckConfig
	GetCheckInterval() *int32
	SetFailureCount(v int32) *HealthCheckConfig
	GetFailureCount() *int32
	SetPort(v int32) *HealthCheckConfig
	GetPort() *int32
	SetProtocol(v string) *HealthCheckConfig
	GetProtocol() *string
	SetTimeoutSeconds(v int32) *HealthCheckConfig
	GetTimeoutSeconds() *int32
}

type HealthCheckConfig struct {
	// The check interval in seconds. Default value: 30 seconds.
	//
	// example:
	//
	// 30
	CheckInterval *int32 `json:"CheckInterval,omitempty" xml:"CheckInterval,omitempty"`
	// The failure count. Default value: 3.
	//
	// example:
	//
	// 3
	FailureCount *int32 `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	// The port number. This parameter is optional.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values: TCP and UDP. This parameter is required only in Continuous mode.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The timeout period for a single check in seconds. Default value: 5 seconds.
	//
	// example:
	//
	// 5
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s HealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s HealthCheckConfig) GoString() string {
	return s.String()
}

func (s *HealthCheckConfig) GetCheckInterval() *int32 {
	return s.CheckInterval
}

func (s *HealthCheckConfig) GetFailureCount() *int32 {
	return s.FailureCount
}

func (s *HealthCheckConfig) GetPort() *int32 {
	return s.Port
}

func (s *HealthCheckConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *HealthCheckConfig) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *HealthCheckConfig) SetCheckInterval(v int32) *HealthCheckConfig {
	s.CheckInterval = &v
	return s
}

func (s *HealthCheckConfig) SetFailureCount(v int32) *HealthCheckConfig {
	s.FailureCount = &v
	return s
}

func (s *HealthCheckConfig) SetPort(v int32) *HealthCheckConfig {
	s.Port = &v
	return s
}

func (s *HealthCheckConfig) SetProtocol(v string) *HealthCheckConfig {
	s.Protocol = &v
	return s
}

func (s *HealthCheckConfig) SetTimeoutSeconds(v int32) *HealthCheckConfig {
	s.TimeoutSeconds = &v
	return s
}

func (s *HealthCheckConfig) Validate() error {
	return dara.Validate(s)
}
