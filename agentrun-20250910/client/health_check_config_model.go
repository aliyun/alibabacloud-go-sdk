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
	// 在将容器视为不健康之前，连续失败的健康检查次数
	FailureThreshold *int32 `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	// 用于健康检查的HTTP GET请求的URL地址
	HttpGetUrl *string `json:"httpGetUrl,omitempty" xml:"httpGetUrl,omitempty"`
	// 在容器启动后，首次执行健康检查前的延迟时间（秒）
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	// 执行健康检查的时间间隔（秒）
	PeriodSeconds *int32 `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	// 在将容器视为健康之前，连续成功的健康检查次数
	SuccessThreshold *int32 `json:"successThreshold,omitempty" xml:"successThreshold,omitempty"`
	// 健康检查的超时时间（秒）
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
