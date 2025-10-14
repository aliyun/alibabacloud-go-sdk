// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmMonitorTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCloudGtmMonitorTemplateRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *CreateCloudGtmMonitorTemplateRequest
	GetClientToken() *string
	SetEvaluationCount(v int32) *CreateCloudGtmMonitorTemplateRequest
	GetEvaluationCount() *int32
	SetExtendInfo(v string) *CreateCloudGtmMonitorTemplateRequest
	GetExtendInfo() *string
	SetFailureRate(v int32) *CreateCloudGtmMonitorTemplateRequest
	GetFailureRate() *int32
	SetInterval(v int32) *CreateCloudGtmMonitorTemplateRequest
	GetInterval() *int32
	SetIpVersion(v string) *CreateCloudGtmMonitorTemplateRequest
	GetIpVersion() *string
	SetIspCityNodes(v []*CreateCloudGtmMonitorTemplateRequestIspCityNodes) *CreateCloudGtmMonitorTemplateRequest
	GetIspCityNodes() []*CreateCloudGtmMonitorTemplateRequestIspCityNodes
	SetName(v string) *CreateCloudGtmMonitorTemplateRequest
	GetName() *string
	SetProtocol(v string) *CreateCloudGtmMonitorTemplateRequest
	GetProtocol() *string
	SetTimeout(v int32) *CreateCloudGtmMonitorTemplateRequest
	GetTimeout() *int32
}

type CreateCloudGtmMonitorTemplateRequest struct {
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
	IspCityNodes []*CreateCloudGtmMonitorTemplateRequestIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Repeated"`
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

func (s CreateCloudGtmMonitorTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmMonitorTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetFailureRate() *int32 {
	return s.FailureRate
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetIspCityNodes() []*CreateCloudGtmMonitorTemplateRequestIspCityNodes {
	return s.IspCityNodes
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateCloudGtmMonitorTemplateRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetAcceptLanguage(v string) *CreateCloudGtmMonitorTemplateRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetClientToken(v string) *CreateCloudGtmMonitorTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetEvaluationCount(v int32) *CreateCloudGtmMonitorTemplateRequest {
	s.EvaluationCount = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetExtendInfo(v string) *CreateCloudGtmMonitorTemplateRequest {
	s.ExtendInfo = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetFailureRate(v int32) *CreateCloudGtmMonitorTemplateRequest {
	s.FailureRate = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetInterval(v int32) *CreateCloudGtmMonitorTemplateRequest {
	s.Interval = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetIpVersion(v string) *CreateCloudGtmMonitorTemplateRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetIspCityNodes(v []*CreateCloudGtmMonitorTemplateRequestIspCityNodes) *CreateCloudGtmMonitorTemplateRequest {
	s.IspCityNodes = v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetName(v string) *CreateCloudGtmMonitorTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetProtocol(v string) *CreateCloudGtmMonitorTemplateRequest {
	s.Protocol = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) SetTimeout(v int32) *CreateCloudGtmMonitorTemplateRequest {
	s.Timeout = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequest) Validate() error {
	if s.IspCityNodes != nil {
		for _, item := range s.IspCityNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCloudGtmMonitorTemplateRequestIspCityNodes struct {
	// The city code of the health check node.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The Internet service provider (ISP) code of the health check node.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s CreateCloudGtmMonitorTemplateRequestIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmMonitorTemplateRequestIspCityNodes) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmMonitorTemplateRequestIspCityNodes) GetCityCode() *string {
	return s.CityCode
}

func (s *CreateCloudGtmMonitorTemplateRequestIspCityNodes) GetIspCode() *string {
	return s.IspCode
}

func (s *CreateCloudGtmMonitorTemplateRequestIspCityNodes) SetCityCode(v string) *CreateCloudGtmMonitorTemplateRequestIspCityNodes {
	s.CityCode = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequestIspCityNodes) SetIspCode(v string) *CreateCloudGtmMonitorTemplateRequestIspCityNodes {
	s.IspCode = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateRequestIspCityNodes) Validate() error {
	return dara.Validate(s)
}
