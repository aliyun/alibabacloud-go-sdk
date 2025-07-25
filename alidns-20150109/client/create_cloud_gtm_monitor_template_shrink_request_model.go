// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmMonitorTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetClientToken() *string
	SetEvaluationCount(v int32) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetEvaluationCount() *int32
	SetExtendInfo(v string) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetExtendInfo() *string
	SetFailureRate(v int32) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetFailureRate() *int32
	SetInterval(v int32) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetInterval() *int32
	SetIpVersion(v string) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetIpVersion() *string
	SetIspCityNodesShrink(v string) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetIspCityNodesShrink() *string
	SetName(v string) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetName() *string
	SetProtocol(v string) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetProtocol() *string
	SetTimeout(v int32) *CreateCloudGtmMonitorTemplateShrinkRequest
	GetTimeout() *int32
}

type CreateCloudGtmMonitorTemplateShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The extended information. The value of this parameter is a JSON string. The required parameters vary based on the health check protocol.
	//
	// 	- HTTP or HTTPS:
	//
	//     **host**: the Host field of an HTTP or HTTPS request header during an HTTP or HTTPS health check. The parameter value indicates the HTTP website that you want to visit. By default, the value is the primary domain name. You can change the value based on your business requirements.
	//
	//     **path**: the URL for HTTP or HTTPS health checks. Default value: /.
	//
	//     **code**: the alert threshold. During an HTTP or HTTPS health check, the system checks whether a web server functions as expected based on the status code that is returned from the web server. If the returned status code is greater than the specified threshold, the corresponding application service address is deemed abnormal. Valid values:
	//
	//     	- 400: specifies an invalid request. If an HTTP or HTTPS request contains invalid request parameters, a web server returns a status code that is greater than 400. You must set path to an exact URL if you set code to 400.
	//
	//     	- 500: specifies a server error. If some exceptions occur on a web server, the web server returns a status code that is greater than 500. This value is used by default.
	//
	//     **sni**: specifies whether to enable Server Name Indication (SNI). This parameter is used only when the health check protocol is HTTPS. SNI is an extension to the Transport Layer Security (TLS) protocol, which allows a client to specify the host to be connected when the client sends a TLS handshake request. TLS handshakes occur before any data of HTTP requests is sent. Therefore, SNI enables servers to identify the services that clients are attempting to access before certificates are sent. This allows the servers to present correct certificates to the clients. Valid values:
	//
	//     	- true: enables SNI.
	//
	//     	- false: disables SNI.
	//
	//     **followRedirect**: specifies whether to follow 3XX redirects. Valid values:
	//
	//     	- true: follows 3XX redirects. You are redirected to the destination address if a 3XX status code such as 301, 302, 303, 307, or 308 is returned.
	//
	//     	- false: does not follow 3XX redirects.
	//
	// 	- ping:
	//
	//     **packetNum**: the total number of Internet Control Message Protocol (ICMP) packets that are sent to the address for each ping-based health check. Valid values: 20, 50, and 100.
	//
	//     **packetLossRate**: the ICMP packet loss rate for each ping-based health check. The packet loss rate in a health check can be calculated by using the following formula: Packet loss rate in a health check = (Number of lost packets/Total number of sent ICMP packets) × 100%. If the packet loss rate reaches the threshold, an alert is triggered. Valid values: 10, 30, 40, 80, 90, and 100.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	FailureRate *int32 `json:"FailureRate,omitempty" xml:"FailureRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The IP address type of health check nodes. Valid values:
	//
	// 	- IPv4: You can set IpVersion to IPv4 to perform health checks on IPv4 addresses.
	//
	// 	- IPv6: You can set IpVersion to IPv6 to perform health checks on IPv6 addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The health check nodes. You can call the [ListCloudGtmMonitorNodes](~~ListCloudGtmMonitorNodes~~) operation to obtain the health check nodes.
	//
	// This parameter is required.
	IspCityNodesShrink *string `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty"`
	// The name of the health check template. We recommend that you use a name that distinguishes the type of health check protocol used.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ping-IPv4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ping
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateCloudGtmMonitorTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmMonitorTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetFailureRate() *int32 {
	return s.FailureRate
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetIspCityNodesShrink() *string {
	return s.IspCityNodesShrink
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetAcceptLanguage(v string) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetClientToken(v string) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetEvaluationCount(v int32) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.EvaluationCount = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetExtendInfo(v string) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.ExtendInfo = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetFailureRate(v int32) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.FailureRate = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetInterval(v int32) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.Interval = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetIpVersion(v string) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetIspCityNodesShrink(v string) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.IspCityNodesShrink = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetName(v string) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetProtocol(v string) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) SetTimeout(v int32) *CreateCloudGtmMonitorTemplateShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
