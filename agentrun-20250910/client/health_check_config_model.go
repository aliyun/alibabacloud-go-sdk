// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHealthCheckConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFailureThreshold(v int32) *HealthCheckConfig
	GetFailureThreshold() *int32
	SetHttpGetUrl(v string) *HealthCheckConfig
	GetHttpGetUrl() *string
	SetInitialDelaySeconds(v int32) *HealthCheckConfig
	GetInitialDelaySeconds() *int32
	SetPeriodSeconds(v int32) *HealthCheckConfig
	GetPeriodSeconds() *int32
	SetSuccessThreshold(v int32) *HealthCheckConfig
	GetSuccessThreshold() *int32
	SetTimeoutSeconds(v int32) *HealthCheckConfig
	GetTimeoutSeconds() *int32
}

type HealthCheckConfig struct {
	// The number of consecutive failed health checks before the container is considered unhealthy
	FailureThreshold *int32 `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	// The URL address for the HTTP GET request used in health checks
	HttpGetUrl *string `json:"httpGetUrl,omitempty" xml:"httpGetUrl,omitempty"`
	// The delay time (in seconds) after the container starts before the first health check is executed
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	// The time interval (in seconds) between health checks
	PeriodSeconds *int32 `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	// The number of consecutive successful health checks required before the container is considered healthy
	SuccessThreshold *int32 `json:"successThreshold,omitempty" xml:"successThreshold,omitempty"`
	// The timeout duration (in seconds) for health checks
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s HealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s HealthCheckConfig) GoString() string {
	return s.String()
}

func (s *HealthCheckConfig) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *HealthCheckConfig) GetHttpGetUrl() *string {
	return s.HttpGetUrl
}

func (s *HealthCheckConfig) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *HealthCheckConfig) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *HealthCheckConfig) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *HealthCheckConfig) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *HealthCheckConfig) SetFailureThreshold(v int32) *HealthCheckConfig {
	s.FailureThreshold = &v
	return s
}

func (s *HealthCheckConfig) SetHttpGetUrl(v string) *HealthCheckConfig {
	s.HttpGetUrl = &v
	return s
}

func (s *HealthCheckConfig) SetInitialDelaySeconds(v int32) *HealthCheckConfig {
	s.InitialDelaySeconds = &v
	return s
}

func (s *HealthCheckConfig) SetPeriodSeconds(v int32) *HealthCheckConfig {
	s.PeriodSeconds = &v
	return s
}

func (s *HealthCheckConfig) SetSuccessThreshold(v int32) *HealthCheckConfig {
	s.SuccessThreshold = &v
	return s
}

func (s *HealthCheckConfig) SetTimeoutSeconds(v int32) *HealthCheckConfig {
	s.TimeoutSeconds = &v
	return s
}

func (s *HealthCheckConfig) Validate() error {
	return dara.Validate(s)
}
