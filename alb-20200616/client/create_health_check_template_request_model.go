// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHealthCheckTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateHealthCheckTemplateRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateHealthCheckTemplateRequest
	GetDryRun() *bool
	SetHealthCheckCodes(v []*string) *CreateHealthCheckTemplateRequest
	GetHealthCheckCodes() []*string
	SetHealthCheckConnectPort(v int32) *CreateHealthCheckTemplateRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckHost(v string) *CreateHealthCheckTemplateRequest
	GetHealthCheckHost() *string
	SetHealthCheckHttpVersion(v string) *CreateHealthCheckTemplateRequest
	GetHealthCheckHttpVersion() *string
	SetHealthCheckInterval(v int32) *CreateHealthCheckTemplateRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *CreateHealthCheckTemplateRequest
	GetHealthCheckMethod() *string
	SetHealthCheckPath(v string) *CreateHealthCheckTemplateRequest
	GetHealthCheckPath() *string
	SetHealthCheckProtocol(v string) *CreateHealthCheckTemplateRequest
	GetHealthCheckProtocol() *string
	SetHealthCheckTemplateName(v string) *CreateHealthCheckTemplateRequest
	GetHealthCheckTemplateName() *string
	SetHealthCheckTimeout(v int32) *CreateHealthCheckTemplateRequest
	GetHealthCheckTimeout() *int32
	SetHealthyThreshold(v int32) *CreateHealthCheckTemplateRequest
	GetHealthyThreshold() *int32
	SetResourceGroupId(v string) *CreateHealthCheckTemplateRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateHealthCheckTemplateRequestTag) *CreateHealthCheckTemplateRequest
	GetTag() []*CreateHealthCheckTemplateRequestTag
	SetUnhealthyThreshold(v int32) *CreateHealthCheckTemplateRequest
	GetUnhealthyThreshold() *int32
}

type CreateHealthCheckTemplateRequest struct {
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
	// The port that is used for health checks.
	//
	// Valid values: **0 to 65535**.
	//
	// Default value: **0**. If you set the value to 0, the port of a backend server is used for health checks.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$SERVER_IP**: the private IP addresses of backend servers. If an IP address is specified, or this parameter is not specified, the ALB instance uses the private IP addresses of backend servers as domain names for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// >  This parameter takes effect only if `HealthCheckProtocol` is set to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// $_ip
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version for health checks.
	//
	// Valid values: **HTTP 1.0*	- and **HTTP 1.1**.
	//
	// Default value: **HTTP 1.1**.
	//
	// >  This parameter is available only if `HealthCheckProtocol` is set to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// HTTP 1.0
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed.
	//
	// Valid values: **1 to 50**.
	//
	// Default value: **2**.
	//
	// example:
	//
	// 2
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// 	- **HEAD*	- (default): By default, HTTP and HTTPS health checks use the HEAD method.
	//
	// 	- **POST**: gRPC health checks use the POST method by default.
	//
	// 	- **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	//
	// >  This parameter is available only if **HealthCheckProtocol*	- is set to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// HEAD
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, the following special characters: - / . % ? # &, and the following extended characters: `_ ; ~ ! ( ) 	- [ ] @ $ ^ : \\" , +`. The URL must start with a forward slash (/).
	//
	// >  This parameter is available only if `HealthCheckProtocol` is set to **HTTP*	- or **HTTPS**.
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
	// The name of the health check template.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// HealthCheckTemplate1
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// The timeout period of a health check response. If a backend server does not respond within the specified timeout period, the backend server is declared unhealthy.
	//
	// Valid values: **1 to 300**. Unit: seconds.
	//
	// Default value: **5**.
	//
	// example:
	//
	// 5
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
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*CreateHealthCheckTemplateRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s CreateHealthCheckTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHealthCheckTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateHealthCheckTemplateRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckCodes() []*string {
	return s.HealthCheckCodes
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckTemplateName() *string {
	return s.HealthCheckTemplateName
}

func (s *CreateHealthCheckTemplateRequest) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *CreateHealthCheckTemplateRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateHealthCheckTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateHealthCheckTemplateRequest) GetTag() []*CreateHealthCheckTemplateRequestTag {
	return s.Tag
}

func (s *CreateHealthCheckTemplateRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateHealthCheckTemplateRequest) SetClientToken(v string) *CreateHealthCheckTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetDryRun(v bool) *CreateHealthCheckTemplateRequest {
	s.DryRun = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckCodes(v []*string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckCodes = v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckConnectPort(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckHost(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckHttpVersion(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckInterval(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckMethod(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckPath(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckProtocol(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckTemplateName(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckTimeout(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthyThreshold(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetResourceGroupId(v string) *CreateHealthCheckTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetTag(v []*CreateHealthCheckTemplateRequestTag) *CreateHealthCheckTemplateRequest {
	s.Tag = v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetUnhealthyThreshold(v int32) *CreateHealthCheckTemplateRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type CreateHealthCheckTemplateRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHealthCheckTemplateRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateHealthCheckTemplateRequestTag) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateHealthCheckTemplateRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateHealthCheckTemplateRequestTag) SetKey(v string) *CreateHealthCheckTemplateRequestTag {
	s.Key = &v
	return s
}

func (s *CreateHealthCheckTemplateRequestTag) SetValue(v string) *CreateHealthCheckTemplateRequestTag {
	s.Value = &v
	return s
}

func (s *CreateHealthCheckTemplateRequestTag) Validate() error {
	return dara.Validate(s)
}
