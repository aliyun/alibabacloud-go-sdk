// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHealthCheckConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetFailureThreshold(v int32) *HealthCheckConfiguration
	GetFailureThreshold() *int32
	SetHttpGetUrl(v string) *HealthCheckConfiguration
	GetHttpGetUrl() *string
	SetInitialDelaySeconds(v int32) *HealthCheckConfiguration
	GetInitialDelaySeconds() *int32
	SetPeriodSeconds(v int32) *HealthCheckConfiguration
	GetPeriodSeconds() *int32
	SetSuccessThreshold(v int32) *HealthCheckConfiguration
	GetSuccessThreshold() *int32
	SetTimeoutSeconds(v int32) *HealthCheckConfiguration
	GetTimeoutSeconds() *int32
}

type HealthCheckConfiguration struct {
	// 在将容器视为不健康之前，连续失败的健康检查次数
	//
	// example:
	//
	// 3
	FailureThreshold *int32 `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	// 用于健康检查的HTTP GET请求的URL地址
	//
	// example:
	//
	// /ready
	HttpGetUrl *string `json:"httpGetUrl,omitempty" xml:"httpGetUrl,omitempty"`
	// 在容器启动后，首次执行健康检查前的延迟时间（秒）
	//
	// example:
	//
	// 30
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	// 执行健康检查的时间间隔（秒）
	//
	// example:
	//
	// 30
	PeriodSeconds *int32 `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	// 在将容器视为健康之前，连续成功的健康检查次数
	//
	// example:
	//
	// 1
	SuccessThreshold *int32 `json:"successThreshold,omitempty" xml:"successThreshold,omitempty"`
	// 健康检查的超时时间（秒）
	//
	// example:
	//
	// 3
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s HealthCheckConfiguration) String() string {
	return dara.Prettify(s)
}

func (s HealthCheckConfiguration) GoString() string {
	return s.String()
}

func (s *HealthCheckConfiguration) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *HealthCheckConfiguration) GetHttpGetUrl() *string {
	return s.HttpGetUrl
}

func (s *HealthCheckConfiguration) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *HealthCheckConfiguration) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *HealthCheckConfiguration) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *HealthCheckConfiguration) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *HealthCheckConfiguration) SetFailureThreshold(v int32) *HealthCheckConfiguration {
	s.FailureThreshold = &v
	return s
}

func (s *HealthCheckConfiguration) SetHttpGetUrl(v string) *HealthCheckConfiguration {
	s.HttpGetUrl = &v
	return s
}

func (s *HealthCheckConfiguration) SetInitialDelaySeconds(v int32) *HealthCheckConfiguration {
	s.InitialDelaySeconds = &v
	return s
}

func (s *HealthCheckConfiguration) SetPeriodSeconds(v int32) *HealthCheckConfiguration {
	s.PeriodSeconds = &v
	return s
}

func (s *HealthCheckConfiguration) SetSuccessThreshold(v int32) *HealthCheckConfiguration {
	s.SuccessThreshold = &v
	return s
}

func (s *HealthCheckConfiguration) SetTimeoutSeconds(v int32) *HealthCheckConfiguration {
	s.TimeoutSeconds = &v
	return s
}

func (s *HealthCheckConfiguration) Validate() error {
	return dara.Validate(s)
}
