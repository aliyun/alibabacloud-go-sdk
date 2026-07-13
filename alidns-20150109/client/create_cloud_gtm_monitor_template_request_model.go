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
	// - zh-CN: Chinese.
	//
	// - en-US: English. This is the default value.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Make sure that the client token is unique for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of consecutive failures that must occur before the system considers the application service unhealthy. This setting helps prevent false alarms caused by transient issues such as network jitter. Valid values:
	//
	// - 1
	//
	// - 2
	//
	// - 3
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The extended information in a JSON string. The parameters vary based on the protocol.
	//
	// - http(s):
	//
	//   **host**: The Host field in the header of the HTTP or HTTPS request. This field identifies the website that you want to access. The default value is the primary domain name. If the destination website uses a specific host, change this value as needed.
	//
	//   **path**: The URL path for the HTTP or HTTPS health check. The default value is "/".
	//
	//   **code**: For an HTTP or HTTPS health check, the system determines whether the web server is working correctly based on the return code. If the return code is greater than this threshold, the system considers the application service unhealthy.
	//
	//   - 400: Bad Request. If an HTTP or HTTPS request contains incorrect parameters, the web server returns a code greater than 400. If you set the threshold to 400, make sure that you specify the exact URL path.
	//
	//   - 500: Server Error. If an exception occurs on the web server, it returns a code greater than 500. The default threshold is 500.
	//
	//   **sni**: Specifies whether to enable Server Name Indication (SNI). This parameter applies only to the HTTPS protocol. SNI is a Transport Layer Security (TLS) extension that allows a client to specify the hostname to connect to at the start of the TLS handshake. This allows the server to present the correct certificate for the requested service.
	//
	//   - true: Enable SNI.
	//
	//   - false: Disable SNI.
	//
	//   **followRedirect**: Specifies whether to follow 3xx redirects.
	//
	//   - true: Follows the redirect if the detection point receives a 3xx status code, such as 301, 302, 303, 307, or 308.
	//
	//   - false: Does not follow the redirect.
	//
	// - ping:
	//
	//   **packetNum**: The number of ICMP packets to send for each ping health check. Valid values: 20, 50, and 100.
	//
	//   **packetLossRate**: The packet loss rate that triggers an alarm. For each ping health check, the system calculates the packet loss rate based on the sent ICMP packets. Packet loss rate = (Number of lost packets / Total number of sent ICMP packets) × 100%. An alarm is triggered if the packet loss rate reaches this threshold. Valid values: 10, 30, 40, 80, 90, and 100.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The failure rate threshold. An endpoint is considered unhealthy if the percentage of unhealthy detection points exceeds this value. Valid values:
	//
	// - 20
	//
	// - 50
	//
	// - 80
	//
	// - 100
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	FailureRate *int32 `json:"FailureRate,omitempty" xml:"FailureRate,omitempty"`
	// The health check interval in seconds. The default value is 60. The minimum interval is 15 seconds, which is available only for Ultimate Edition instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The IP address type for health checks.
	//
	// - IPv4: The destination address is an IPv4 address.
	//
	// - IPv6: The destination address is an IPv6 address.
	//
	// This parameter is required.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// A list of detection points. For more information, see [ListCloudGtmMonitorNodes](https://help.aliyun.com/document_detail/2797349.html).
	//
	// This parameter is required.
	IspCityNodes []*CreateCloudGtmMonitorTemplateRequestIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Repeated"`
	// The name of the health check template. Name the template to easily identify the health check protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ping-IPv4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protocol for health checks on the destination IP address.
	//
	// - ping
	//
	// - tcp
	//
	// - http
	//
	// - https
	//
	// This parameter is required.
	//
	// example:
	//
	// ping
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The health check timeout in milliseconds. If a packet is not returned within the timeout period, the health check is considered to have timed out. Valid values:
	//
	// - 2000
	//
	// - 3000
	//
	// - 5000
	//
	// - 10000
	//
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
	// The city code for the detection point.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The ISP code for the detection point.
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
