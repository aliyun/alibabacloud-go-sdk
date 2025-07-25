// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmMonitorTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmMonitorTemplateRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *UpdateCloudGtmMonitorTemplateRequest
	GetClientToken() *string
	SetEvaluationCount(v int32) *UpdateCloudGtmMonitorTemplateRequest
	GetEvaluationCount() *int32
	SetExtendInfo(v string) *UpdateCloudGtmMonitorTemplateRequest
	GetExtendInfo() *string
	SetFailureRate(v int32) *UpdateCloudGtmMonitorTemplateRequest
	GetFailureRate() *int32
	SetInterval(v int32) *UpdateCloudGtmMonitorTemplateRequest
	GetInterval() *int32
	SetIspCityNodes(v []*UpdateCloudGtmMonitorTemplateRequestIspCityNodes) *UpdateCloudGtmMonitorTemplateRequest
	GetIspCityNodes() []*UpdateCloudGtmMonitorTemplateRequestIspCityNodes
	SetName(v string) *UpdateCloudGtmMonitorTemplateRequest
	GetName() *string
	SetTemplateId(v string) *UpdateCloudGtmMonitorTemplateRequest
	GetTemplateId() *string
	SetTimeout(v int32) *UpdateCloudGtmMonitorTemplateRequest
	GetTimeout() *int32
}

type UpdateCloudGtmMonitorTemplateRequest struct {
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of retries. The system will only judge the application service as abnormal after consecutive monitoring failures to prevent inaccurate monitoring results due to momentary network fluctuations or other reasons. Available retry options are:
	//
	// - 1
	//
	// - 2
	//
	// - 3
	//
	// example:
	//
	// 2
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The extended information. The value of this parameter is a JSON string. The required parameters vary based on the health check protocol.
	//
	// 	- HTTP or HTTPS:
	//
	//     host: the Host field of an HTTP or HTTPS request header during an HTTP or HTTPS health check. The parameter value indicates the HTTP website that you want to visit. By default, the value is the primary domain name. You can change the value based on your business requirements.
	//
	//     path: the URL for HTTP or HTTPS health checks. Default value: /.
	//
	//     code: the alert threshold. During an HTTP or HTTPS health check, the system checks whether a web server functions as expected based on the status code that is returned from the web server. If the returned status code is greater than the specified threshold, the corresponding application service address is deemed abnormal. Valid values:
	//
	//     	- 400: specifies an invalid request. If an HTTP or HTTPS request contains invalid request parameters, a web server returns a status code that is greater than 400. You must set path to an exact URL if you set code to 400.
	//
	//     	- 500: specifies a server error. If some exceptions occur on a web server, the web server returns a status code that is greater than 500. This value is used by default.
	//
	//     sni: specifies whether to enable Server Name Indication (SNI). This parameter is used only when the health check protocol is HTTPS. SNI is an extension to the Transport Layer Security (TLS) protocol, which allows a client to specify the host to be connected when the client sends a TLS handshake request. TLS handshakes occur before any data of HTTP requests is sent. Therefore, SNI enables servers to identify the services that clients are attempting to access before certificates are sent. This allows the servers to present correct certificates to the clients. Valid values:
	//
	//     	- true: enables SNI.
	//
	//     	- false: disables SNI.
	//
	//     followRedirect: specifies whether to follow 3XX redirects. Valid values:
	//
	//     	- true: follows 3XX redirects. You are redirected to the destination address if a 3XX status code such as 301, 302, 303, 307, or 308 is returned.
	//
	//     	- false: does not follow 3XX redirects.
	//
	// 	- ping:
	//
	//     packetNum: the total number of Internet Control Message Protocol (ICMP) packets that are sent to the address for each ping-based health check. Valid values: 20, 50, and 100.
	//
	//     packetLossRate: the ICMP packet loss rate for each ping-based health check. The packet loss rate in a health check can be calculated by using the following formula: Packet loss rate in a health check = (Number of lost packets/Total number of sent ICMP packets) × 100%. If the packet loss rate reaches the threshold, an alert is triggered. Valid values: 10, 30, 40, 80, 90, and 100.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// Percentage of selected node probe failures (%), that is, the percentage of abnormal detection points among the total detection points. When the failure ratio exceeds the set threshold, the service address is judged as abnormal. The available failure ratio thresholds are:
	//
	// - 20
	//
	// - 50
	//
	// - 80
	//
	// - 100
	//
	// example:
	//
	// 50
	FailureRate *int32 `json:"FailureRate,omitempty" xml:"FailureRate,omitempty"`
	// The time interval (in seconds) for each health check probe. By default, it probes every 60 seconds. The minimum supported interval for health checks is 15 seconds, available for flagship edition instances.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The health check nodes. You can call the [ListCloudGtmMonitorNodes](https://help.aliyun.com/document_detail/2797327.html) operation to obtain the health check nodes.
	IspCityNodes []*UpdateCloudGtmMonitorTemplateRequestIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Repeated"`
	// The name of the health check probe template, which is generally recommended to be distinguishable and memorable for configuration personnel, ideally indicating the health check protocol for easier identification.
	//
	// example:
	//
	// Ping-IPv4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the health check template that you want to modify. This ID uniquely identifies the health check template.
	//
	// This parameter is required.
	//
	// example:
	//
	// mtp-89518052425100**80
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Probe timeout (in milliseconds), data packets not returned within the timeout period are considered as health check timeouts:
	//
	// - 2000
	//
	// - 3000
	//
	// - 5000
	//
	// - 10000
	//
	// example:
	//
	// 5000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateCloudGtmMonitorTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmMonitorTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetFailureRate() *int32 {
	return s.FailureRate
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetIspCityNodes() []*UpdateCloudGtmMonitorTemplateRequestIspCityNodes {
	return s.IspCityNodes
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateCloudGtmMonitorTemplateRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetAcceptLanguage(v string) *UpdateCloudGtmMonitorTemplateRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetClientToken(v string) *UpdateCloudGtmMonitorTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetEvaluationCount(v int32) *UpdateCloudGtmMonitorTemplateRequest {
	s.EvaluationCount = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetExtendInfo(v string) *UpdateCloudGtmMonitorTemplateRequest {
	s.ExtendInfo = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetFailureRate(v int32) *UpdateCloudGtmMonitorTemplateRequest {
	s.FailureRate = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetInterval(v int32) *UpdateCloudGtmMonitorTemplateRequest {
	s.Interval = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetIspCityNodes(v []*UpdateCloudGtmMonitorTemplateRequestIspCityNodes) *UpdateCloudGtmMonitorTemplateRequest {
	s.IspCityNodes = v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetName(v string) *UpdateCloudGtmMonitorTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetTemplateId(v string) *UpdateCloudGtmMonitorTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) SetTimeout(v int32) *UpdateCloudGtmMonitorTemplateRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudGtmMonitorTemplateRequestIspCityNodes struct {
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

func (s UpdateCloudGtmMonitorTemplateRequestIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmMonitorTemplateRequestIspCityNodes) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmMonitorTemplateRequestIspCityNodes) GetCityCode() *string {
	return s.CityCode
}

func (s *UpdateCloudGtmMonitorTemplateRequestIspCityNodes) GetIspCode() *string {
	return s.IspCode
}

func (s *UpdateCloudGtmMonitorTemplateRequestIspCityNodes) SetCityCode(v string) *UpdateCloudGtmMonitorTemplateRequestIspCityNodes {
	s.CityCode = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequestIspCityNodes) SetIspCode(v string) *UpdateCloudGtmMonitorTemplateRequestIspCityNodes {
	s.IspCode = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRequestIspCityNodes) Validate() error {
	return dara.Validate(s)
}
