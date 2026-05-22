// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckFailTimes(v int32) *CreateEdgeContainerAppRequest
	GetHealthCheckFailTimes() *int32
	SetHealthCheckHost(v string) *CreateEdgeContainerAppRequest
	GetHealthCheckHost() *string
	SetHealthCheckHttpCode(v string) *CreateEdgeContainerAppRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *CreateEdgeContainerAppRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *CreateEdgeContainerAppRequest
	GetHealthCheckMethod() *string
	SetHealthCheckPort(v int32) *CreateEdgeContainerAppRequest
	GetHealthCheckPort() *int32
	SetHealthCheckSuccTimes(v int32) *CreateEdgeContainerAppRequest
	GetHealthCheckSuccTimes() *int32
	SetHealthCheckTimeout(v int32) *CreateEdgeContainerAppRequest
	GetHealthCheckTimeout() *int32
	SetHealthCheckType(v string) *CreateEdgeContainerAppRequest
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *CreateEdgeContainerAppRequest
	GetHealthCheckURI() *string
	SetName(v string) *CreateEdgeContainerAppRequest
	GetName() *string
	SetRemarks(v string) *CreateEdgeContainerAppRequest
	GetRemarks() *string
	SetServicePort(v int32) *CreateEdgeContainerAppRequest
	GetServicePort() *int32
	SetTargetPort(v int32) *CreateEdgeContainerAppRequest
	GetTargetPort() *int32
}

type CreateEdgeContainerAppRequest struct {
	HealthCheckFailTimes *int32  `json:"HealthCheckFailTimes,omitempty" xml:"HealthCheckFailTimes,omitempty"`
	HealthCheckHost      *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	HealthCheckHttpCode  *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckInterval  *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckMethod    *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	HealthCheckPort      *int32  `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckSuccTimes *int32  `json:"HealthCheckSuccTimes,omitempty" xml:"HealthCheckSuccTimes,omitempty"`
	HealthCheckTimeout   *int32  `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	HealthCheckType      *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI       *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// This parameter is required.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// This parameter is required.
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// This parameter is required.
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s CreateEdgeContainerAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckFailTimes() *int32 {
	return s.HealthCheckFailTimes
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckPort() *int32 {
	return s.HealthCheckPort
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckSuccTimes() *int32 {
	return s.HealthCheckSuccTimes
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *CreateEdgeContainerAppRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *CreateEdgeContainerAppRequest) GetName() *string {
	return s.Name
}

func (s *CreateEdgeContainerAppRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *CreateEdgeContainerAppRequest) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *CreateEdgeContainerAppRequest) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckFailTimes(v int32) *CreateEdgeContainerAppRequest {
	s.HealthCheckFailTimes = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckHost(v string) *CreateEdgeContainerAppRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckHttpCode(v string) *CreateEdgeContainerAppRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckInterval(v int32) *CreateEdgeContainerAppRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckMethod(v string) *CreateEdgeContainerAppRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckPort(v int32) *CreateEdgeContainerAppRequest {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckSuccTimes(v int32) *CreateEdgeContainerAppRequest {
	s.HealthCheckSuccTimes = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckTimeout(v int32) *CreateEdgeContainerAppRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckType(v string) *CreateEdgeContainerAppRequest {
	s.HealthCheckType = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetHealthCheckURI(v string) *CreateEdgeContainerAppRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetName(v string) *CreateEdgeContainerAppRequest {
	s.Name = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetRemarks(v string) *CreateEdgeContainerAppRequest {
	s.Remarks = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetServicePort(v int32) *CreateEdgeContainerAppRequest {
	s.ServicePort = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) SetTargetPort(v int32) *CreateEdgeContainerAppRequest {
	s.TargetPort = &v
	return s
}

func (s *CreateEdgeContainerAppRequest) Validate() error {
	return dara.Validate(s)
}
