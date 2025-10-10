// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHealthCheckTemplateAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateHealthCheckTemplateAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateHealthCheckTemplateAttributeRequest
	GetDryRun() *bool
	SetHealthCheckCodes(v []*string) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckCodes() []*string
	SetHealthCheckConnectPort(v int32) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckHost(v string) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckHost() *string
	SetHealthCheckHttpVersion(v string) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckHttpVersion() *string
	SetHealthCheckInterval(v int32) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckMethod() *string
	SetHealthCheckPath(v string) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckPath() *string
	SetHealthCheckProtocol(v string) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckProtocol() *string
	SetHealthCheckTemplateId(v string) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckTemplateId() *string
	SetHealthCheckTemplateName(v string) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckTemplateName() *string
	SetHealthCheckTimeout(v int32) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthCheckTimeout() *int32
	SetHealthyThreshold(v int32) *UpdateHealthCheckTemplateAttributeRequest
	GetHealthyThreshold() *int32
	SetUnhealthyThreshold(v int32) *UpdateHealthCheckTemplateAttributeRequest
	GetUnhealthyThreshold() *int32
}

type UpdateHealthCheckTemplateAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a **2xx*	- HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The HTTP status codes that indicate a healthy backend server.
	//
	// example:
	//
	// 5
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The port that is used for health checks. Valid values: **0 to 65535**. Default value: **0**. This value indicates that the port of a backend server is used for health checks.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$SERVER_IP*	- (default): the private IP address of a backend server. If an IP address is specified, or this parameter is not specified, the ALB instance uses the private IP addresses of backend servers as domain names for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// >  This parameter is available only if `HealthCheckProtocol` is set to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// $_ip
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version that is used for health checks.
	//
	// Valid values: **HTTP1.0*	- and **HTTP1.1**.
	//
	// Default value: **HTTP1.1**.
	//
	// >  This parameter is available only if you set `HealthCheckProtocol` to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// HTTP1.0
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds. Valid values: **1 to 50**. Default value: **2**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// 	- **HEAD*	- (default): By default, HTTP and HTTPS health checks use the HEAD method.
	//
	// 	- **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	//
	// 	- **POST**: gRPC health checks use the POST method by default.
	//
	// >  This parameter is available only if you set **HealthCheckProtocol*	- to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// HEAD
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), ampersands (&), and the following extended character sets: `_ ; ~ ! ( ) 	- [ ] @ $ ^ : \\" , +`.
	//
	// The URL must start with a forward slash (/).
	//
	// >  This parameter is available only if you set **HealthCheckProtocol*	- to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// 	- **HTTP*	- (default): HTTP health checks simulate browser behaviors by sending HEAD or GET requests to probe the availability of backend servers.
	//
	// 	- **HTTPS**: The ALB instance sends HEAD or GET requests, which simulate browser requests, to check whether the backend server is healthy. HTTPS supports encryption and provides higher security than HTTP.
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to check whether the port of the backend server is reachable.
	//
	// 	- **gRPC**: gRPC health checks send POST or GET requests to a backend server to check whether the backend server is healthy.
	//
	// example:
	//
	// HTTP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hct-bp1qjwo61pqz3ahltv0mw
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The name of the health check template.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// HealthCheckTemplate1
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// The timeout period of a health check response. If a backend server does not respond within the specified timeout period, the backend server is declared unhealthy.
	//
	// Unit: seconds. Valid values: **1 to 300**. Default value: **5**.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**.
	//
	// Valid values: **2 to 10**.
	//
	// Default value: **3**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// Valid values: **2 to 10**.
	//
	// Default value: **3**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateHealthCheckTemplateAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckCodes() []*string {
	return s.HealthCheckCodes
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckTemplateId() *string {
	return s.HealthCheckTemplateId
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckTemplateName() *string {
	return s.HealthCheckTemplateName
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UpdateHealthCheckTemplateAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetClientToken(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetDryRun(v bool) *UpdateHealthCheckTemplateAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckCodes(v []*string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckCodes = v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckConnectPort(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckHost(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckHttpVersion(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckInterval(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckMethod(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckPath(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckProtocol(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateId(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateName(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTimeout(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthyThreshold(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetUnhealthyThreshold(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) Validate() error {
	return dara.Validate(s)
}
