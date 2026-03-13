// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomHealthCheckConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFailureThreshold(v int32) *CustomHealthCheckConfig
	GetFailureThreshold() *int32
	SetHttpGetUrl(v string) *CustomHealthCheckConfig
	GetHttpGetUrl() *string
	SetInitialDelaySeconds(v int32) *CustomHealthCheckConfig
	GetInitialDelaySeconds() *int32
	SetPeriodSeconds(v int32) *CustomHealthCheckConfig
	GetPeriodSeconds() *int32
	SetSuccessThreshold(v int32) *CustomHealthCheckConfig
	GetSuccessThreshold() *int32
	SetTimeoutSeconds(v int32) *CustomHealthCheckConfig
	GetTimeoutSeconds() *int32
}

type CustomHealthCheckConfig struct {
	// The threshold for health check failures. When this value is reached, the system considers the health check failed. Valid values: 1 to 120. Default value: 3.
	//
	// example:
	//
	// 1
	FailureThreshold *int32 `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	// The health check URL of the custom container. The URL can be up to 2,048 characters in length.
	//
	// example:
	//
	// /ready
	HttpGetUrl *string `json:"httpGetUrl,omitempty" xml:"httpGetUrl,omitempty"`
	// The delay between the container startup and the health check. Valid values: 0 to 120. Default value: 0.
	//
	// example:
	//
	// 1
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	// The health check period. Valid values: 1 to 120. Default value: 3.
	//
	// example:
	//
	// 1
	PeriodSeconds *int32 `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	// The threshold for health check successes. When this value is reached, the system considers the health check successful. Valid values: 1 to 120. Default value: 1.
	//
	// example:
	//
	// 2
	SuccessThreshold *int32 `json:"successThreshold,omitempty" xml:"successThreshold,omitempty"`
	// The timeout period of the health check. Unit: seconds. Valid values: 1 to 3. Default value: 1.
	//
	// example:
	//
	// 2
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s CustomHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s CustomHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *CustomHealthCheckConfig) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *CustomHealthCheckConfig) GetHttpGetUrl() *string {
	return s.HttpGetUrl
}

func (s *CustomHealthCheckConfig) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *CustomHealthCheckConfig) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *CustomHealthCheckConfig) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *CustomHealthCheckConfig) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CustomHealthCheckConfig) SetFailureThreshold(v int32) *CustomHealthCheckConfig {
	s.FailureThreshold = &v
	return s
}

func (s *CustomHealthCheckConfig) SetHttpGetUrl(v string) *CustomHealthCheckConfig {
	s.HttpGetUrl = &v
	return s
}

func (s *CustomHealthCheckConfig) SetInitialDelaySeconds(v int32) *CustomHealthCheckConfig {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CustomHealthCheckConfig) SetPeriodSeconds(v int32) *CustomHealthCheckConfig {
	s.PeriodSeconds = &v
	return s
}

func (s *CustomHealthCheckConfig) SetSuccessThreshold(v int32) *CustomHealthCheckConfig {
	s.SuccessThreshold = &v
	return s
}

func (s *CustomHealthCheckConfig) SetTimeoutSeconds(v int32) *CustomHealthCheckConfig {
	s.TimeoutSeconds = &v
	return s
}

func (s *CustomHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}
