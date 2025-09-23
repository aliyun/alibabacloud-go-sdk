// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigHealthCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardProtocol(v string) *ConfigHealthCheckRequest
	GetForwardProtocol() *string
	SetFrontendPort(v int32) *ConfigHealthCheckRequest
	GetFrontendPort() *int32
	SetHealthCheck(v string) *ConfigHealthCheckRequest
	GetHealthCheck() *string
	SetInstanceId(v string) *ConfigHealthCheckRequest
	GetInstanceId() *string
}

type ConfigHealthCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tcp
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Type":"tcp","Timeout":10,"Port":80,"Interval":10,"Up":10,"Down":40}"}
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ConfigHealthCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *ConfigHealthCheckRequest) GetForwardProtocol() *string {
	return s.ForwardProtocol
}

func (s *ConfigHealthCheckRequest) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *ConfigHealthCheckRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *ConfigHealthCheckRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigHealthCheckRequest) SetForwardProtocol(v string) *ConfigHealthCheckRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetFrontendPort(v int32) *ConfigHealthCheckRequest {
	s.FrontendPort = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetHealthCheck(v string) *ConfigHealthCheckRequest {
	s.HealthCheck = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetInstanceId(v string) *ConfigHealthCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigHealthCheckRequest) Validate() error {
	return dara.Validate(s)
}
