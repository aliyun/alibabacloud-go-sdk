// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHealthCheckTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHealthCheckTemplates(v []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) *ListHealthCheckTemplatesResponseBody
	GetHealthCheckTemplates() []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates
	SetMaxResults(v int32) *ListHealthCheckTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListHealthCheckTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHealthCheckTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHealthCheckTemplatesResponseBody
	GetTotalCount() *int32
}

type ListHealthCheckTemplatesResponseBody struct {
	// The health check templates.
	HealthCheckTemplates []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates `json:"HealthCheckTemplates,omitempty" xml:"HealthCheckTemplates,omitempty" type:"Repeated"`
	// The number of entries returned per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHealthCheckTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHealthCheckTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponseBody) GetHealthCheckTemplates() []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	return s.HealthCheckTemplates
}

func (s *ListHealthCheckTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHealthCheckTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHealthCheckTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHealthCheckTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHealthCheckTemplatesResponseBody) SetHealthCheckTemplates(v []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) *ListHealthCheckTemplatesResponseBody {
	s.HealthCheckTemplates = v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetMaxResults(v int32) *ListHealthCheckTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetNextToken(v string) *ListHealthCheckTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetRequestId(v string) *ListHealthCheckTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetTotalCount(v int32) *ListHealthCheckTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHealthCheckTemplatesResponseBodyHealthCheckTemplates struct {
	// The HTTP status codes that indicate healthy backend servers.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The port that is used for health checks.
	//
	// Valid values: \\*\\	- 0 to 65535\\*\\*.
	//
	// The default value is **0**, which specifies that the port of a backend server is used for health checks.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$SERVER_IP*	- (default): the private IP address of a backend server. If an IP address is specified, or this parameter is not specified, the ALB instance uses the private IP address of each backend server as the domain name for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// >  This parameter takes effect only if you set `HealthCheckProtocol` to **HTTP*	- or **HTTPS**.
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
	// >  This parameter takes effect only if you set `HealthCheckProtocol` to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// HTTP 1.0
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
	// >  This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// HEAD
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL path that you want to use for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, the following special characters: - / . % ? # &, and the following extended characters: `_ ; ~ ! ( ) 	- [ ] @ $ ^ : \\" , +`. The URL must start with a forward slash (/).
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// 	- **HTTP*	- (default): The ALB instance sends HEAD or GET requests, which simulate browser requests, to check whether the backend server is healthy.
	//
	// 	- **HTTPS**: HTTPS health checks simulate browser behaviors by sending HEAD or GET requests to probe the availability of backend servers. HTTPS provides higher security because HTTPS supports data encryption.
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to check whether the port of the backend server is reachable.
	//
	// 	- **gRPC**: gRPC health checks send POST or GET requests to a backend server to check whether the backend server is healthy.
	//
	// example:
	//
	// HTTP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The ID of the health check template.
	//
	// example:
	//
	// hct-bp1qjwo61pqz3ahltv****
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// The name of the health check template.
	//
	// The name must be 2 to 128 character characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// HealthCheckTemplate1
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// The timeout period of a health check response. If a backend Elastic Compute Service (ECS) instance does not respond within the specified timeout period, the ECS instance fails to pass the health check.
	//
	// Valid values: **1 to 300**. Unit: seconds.
	//
	// Default value: **5**.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status changes from **fail*	- to **success**.
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
	Tags []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status changes from **success*	- to **fail**.
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

func (s ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckCodes() []*string {
	return s.HealthCheckCodes
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckTemplateId() *string {
	return s.HealthCheckTemplateId
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckTemplateName() *string {
	return s.HealthCheckTemplateName
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetTags() []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags {
	return s.Tags
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckCodes(v []*string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckCodes = v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckConnectPort(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckHost(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckHost = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckHttpVersion(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckInterval(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckMethod(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckMethod = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckPath(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckPath = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckProtocol(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTemplateId(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTemplateName(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTimeout(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTimeout = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthyThreshold(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthyThreshold = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetResourceGroupId(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.ResourceGroupId = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetTags(v []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.Tags = v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetUnhealthyThreshold(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.UnhealthyThreshold = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) Validate() error {
	return dara.Validate(s)
}

type ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags struct {
	// The tag key. The tag key can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags) String() string {
	return dara.Prettify(s)
}

func (s ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags) GetKey() *string {
	return s.Key
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags) GetValue() *string {
	return s.Value
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags) SetKey(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags {
	s.Key = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags) SetValue(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags {
	s.Value = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplatesTags) Validate() error {
	return dara.Validate(s)
}
