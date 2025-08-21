// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHealthCheck interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheck(v string) *HealthCheck
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *HealthCheck
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *HealthCheck
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckDomain(v string) *HealthCheck
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *HealthCheck
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *HealthCheck
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *HealthCheck
	GetHealthCheckMethod() *string
	SetHealthCheckTimeout(v int32) *HealthCheck
	GetHealthCheckTimeout() *int32
	SetHealthCheckType(v string) *HealthCheck
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *HealthCheck
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *HealthCheck
	GetHealthyThreshold() *int32
	SetUnhealthyThreshold(v int32) *HealthCheck
	GetUnhealthyThreshold() *int32
}

type HealthCheck struct {
	HealthCheck               *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckConnectPort    *int32  `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	HealthCheckConnectTimeout *int32  `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	HealthCheckDomain         *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode       *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval       *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod         *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckTimeout        *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckType           *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI            *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold          *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	UnhealthyThreshold        *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s HealthCheck) String() string {
	return dara.Prettify(s)
}

func (s HealthCheck) GoString() string {
	return s.String()
}

func (s *HealthCheck) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *HealthCheck) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *HealthCheck) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *HealthCheck) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *HealthCheck) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *HealthCheck) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *HealthCheck) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *HealthCheck) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *HealthCheck) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *HealthCheck) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *HealthCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *HealthCheck) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *HealthCheck) SetHealthCheck(v string) *HealthCheck {
	s.HealthCheck = &v
	return s
}

func (s *HealthCheck) SetHealthCheckConnectPort(v int32) *HealthCheck {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *HealthCheck) SetHealthCheckConnectTimeout(v int32) *HealthCheck {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *HealthCheck) SetHealthCheckDomain(v string) *HealthCheck {
	s.HealthCheckDomain = &v
	return s
}

func (s *HealthCheck) SetHealthCheckHttpCode(v string) *HealthCheck {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *HealthCheck) SetHealthCheckInterval(v int32) *HealthCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *HealthCheck) SetHealthCheckMethod(v string) *HealthCheck {
	s.HealthCheckMethod = &v
	return s
}

func (s *HealthCheck) SetHealthCheckTimeout(v int32) *HealthCheck {
	s.HealthCheckTimeout = &v
	return s
}

func (s *HealthCheck) SetHealthCheckType(v string) *HealthCheck {
	s.HealthCheckType = &v
	return s
}

func (s *HealthCheck) SetHealthCheckURI(v string) *HealthCheck {
	s.HealthCheckURI = &v
	return s
}

func (s *HealthCheck) SetHealthyThreshold(v int32) *HealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *HealthCheck) SetUnhealthyThreshold(v int32) *HealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

func (s *HealthCheck) Validate() error {
	return dara.Validate(s)
}
