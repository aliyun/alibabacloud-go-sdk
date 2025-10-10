// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHealthCheckTemplateAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckCodes(v []*string) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckCodes() []*string
	SetHealthCheckConnectPort(v int32) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckHost(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckHost() *string
	SetHealthCheckHttpVersion(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckHttpVersion() *string
	SetHealthCheckInterval(v int32) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckMethod() *string
	SetHealthCheckPath(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckPath() *string
	SetHealthCheckProtocol(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckProtocol() *string
	SetHealthCheckTemplateId(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckTemplateId() *string
	SetHealthCheckTemplateName(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckTemplateName() *string
	SetHealthCheckTimeout(v int32) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthCheckTimeout() *int32
	SetHealthyThreshold(v int32) *GetHealthCheckTemplateAttributeResponseBody
	GetHealthyThreshold() *int32
	SetRequestId(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetHealthCheckTemplateAttributeResponseBody
	GetResourceGroupId() *string
	SetTags(v []*GetHealthCheckTemplateAttributeResponseBodyTags) *GetHealthCheckTemplateAttributeResponseBody
	GetTags() []*GetHealthCheckTemplateAttributeResponseBodyTags
	SetUnhealthyThreshold(v int32) *GetHealthCheckTemplateAttributeResponseBody
	GetUnhealthyThreshold() *int32
}

type GetHealthCheckTemplateAttributeResponseBody struct {
	// The HTTP status codes that indicate a healthy backend server.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The port that is used for health checks.
	//
	// Valid values: **0*	- to **65535**.
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
	// $SERVER_IP
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version for health checks.
	//
	// Valid values: **HTTP1.0*	- and **HTTP1.1**.
	//
	// >  This parameter takes effect only if you set `HealthCheckProtocol` to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// HTTP1.0
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds. Valid values: **1 to 50**.
	//
	// example:
	//
	// 3
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// 	- **HEAD*	- (default): By default, HTTP and HTTPS health checks use the HEAD method.
	//
	// 	- **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	//
	// 	- **POST**: gRPC health checks use the POST method by default.
	//
	// >  This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// GET
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, the following special characters: - / . % ? # &, and the following extended characters: `_ ; ~ ! ( ) 	- [ ] @ $ ^ : \\" , +`. The URL must start with a forward slash (/).
	//
	// >  This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**, **HTTPS**, or **gRPC**.
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
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to probe the availability of backend servers.
	//
	// 	- **gRPC**: gRPC health checks send POST or GET requests to a backend server to probe the availability of backend servers.
	//
	// example:
	//
	// HTTP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The ID of the health check template.
	//
	// example:
	//
	// hct-x4jazoyi6tvsq9****
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The name of the health check template.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// HealthCheckTemplate1
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// The timeout period of a health check response. If a backend server does not respond within the specified timeout period, the backend server is declared unhealthy. Unit: seconds.
	//
	// Valid values: **1*	- to **300**.
	//
	// example:
	//
	// 200
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 5
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB1AFC33-DAE8-528E-AA4D-4A6AABE71945
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*GetHealthCheckTemplateAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 5
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s GetHealthCheckTemplateAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckCodes() []*string {
	return s.HealthCheckCodes
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckTemplateId() *string {
	return s.HealthCheckTemplateId
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckTemplateName() *string {
	return s.HealthCheckTemplateName
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetTags() []*GetHealthCheckTemplateAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetHealthCheckTemplateAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckCodes(v []*string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckCodes = v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckConnectPort(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckHost(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckHost = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckHttpVersion(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckInterval(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckMethod(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckPath(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckPath = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckProtocol(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckProtocol = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTemplateId(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTemplateName(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTimeout(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthyThreshold(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetRequestId(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetResourceGroupId(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetTags(v []*GetHealthCheckTemplateAttributeResponseBodyTags) *GetHealthCheckTemplateAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetUnhealthyThreshold(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHealthCheckTemplateAttributeResponseBodyTags struct {
	// The tag key. The tag key can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length, and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetHealthCheckTemplateAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetHealthCheckTemplateAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetHealthCheckTemplateAttributeResponseBodyTags) SetKey(v string) *GetHealthCheckTemplateAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBodyTags) SetValue(v string) *GetHealthCheckTemplateAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
