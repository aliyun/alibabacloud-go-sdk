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
	// The number of consecutive failed health checks required for an application to be considered as unhealthy. Valid values: 1 to 10. Default value: 5.
	//
	// example:
	//
	// 3
	HealthCheckFailTimes *int32 `json:"HealthCheckFailTimes,omitempty" xml:"HealthCheckFailTimes,omitempty"`
	// The domain name that is used for health checks. This parameter is empty by default.
	//
	// example:
	//
	// www.aliyun.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP status code returned for a successful health check. Valid values:
	//
	// 	- **http_2xx*	- (default)
	//
	// 	- **http_3xx**
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds. Valid values: **1*	- to **50**. Default value: **5**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP request method for health checks. Valid values:
	//
	// 	- **HEAD*	- (default): requests the headers of the resource.
	//
	// 	- **GET**: requests the specified resource and returns both the headers and entity body.
	//
	// example:
	//
	// HEAD
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The port used for health checks. Valid values: 1 to 65535. Default value: 80.
	//
	// example:
	//
	// 80
	HealthCheckPort *int32 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The number of consecutive successful health checks required for an application to be considered as healthy. Valid values: 1 to 10. Default value: 2.
	//
	// example:
	//
	// 2
	HealthCheckSuccTimes *int32 `json:"HealthCheckSuccTimes,omitempty" xml:"HealthCheckSuccTimes,omitempty"`
	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. Unit: seconds.\\
	//
	// Valid values: **1*	- to **100**.\\
	//
	// Default value: **3**.
	//
	// example:
	//
	// 5
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The health check type. By default, this parameter is left empty.
	//
	// Valid values:
	//
	// 	- **l4**: Layer 4 health check.
	//
	// 	- **l7**: Layer 7 health check.
	//
	// example:
	//
	// l7
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URI used for health checks. The URI must be **1*	- to **80*	- characters in length. Default value: "/".
	//
	// example:
	//
	// /health_check
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The name of the application. The name must start with a lowercase letter and can contain lowercase letters, digits, and hyphens (-). The name must be 6 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks. This parameter is empty by default.
	//
	// example:
	//
	// test app
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The server port. Valid values: 1 to 65535.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The backend port, which is also the service port of the application. Valid values: 1 to 65535.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
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
