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
	// - zh-CN: Chinese
	//
	// - en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. Make sure that the token is unique for each request. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of consecutive health check failures that must occur before an application service is considered abnormal. This helps prevent false alarms caused by transient issues such as network jitter. Valid values:
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
	// The extended information in a JSON string. The parameters vary based on the health check protocol.
	//
	// - HTTP and HTTPS:
	//
	//   host: When you perform an HTTP or HTTPS health check, this parameter specifies the Host field in the HTTP request header. It identifies the target website. The default value is the primary domain name. If the target website has specific requirements for the Host field, modify this parameter as needed.
	//
	//   path: The path for the HTTP or HTTPS health check. The default value is /.
	//
	//   code: When you perform an HTTP or HTTPS health check, the system uses the return code from the web server to determine its status. If the return code exceeds the specified threshold, the application service is considered abnormal.
	//
	//   - 400: Bad Request. If an HTTP or HTTPS request contains invalid parameters, the web server returns a code greater than 400. If you set the threshold to 400, make sure that you specify the exact URL path.
	//
	//   - 500: Server Error. If the web server encounters an error, it returns a code greater than 500. The default threshold is 500.
	//
	//   sni: Specifies whether to enable Server Name Indication (SNI). This parameter is used only for HTTPS health checks. SNI is an extension to the Transport Layer Security (TLS) protocol. It allows a client to specify the hostname it is trying to connect to at the start of the TLS handshake. Because the TLS handshake occurs before any HTTP data is sent, SNI allows the server to know which service the client is trying to access before sending the certificate. This enables the server to present the correct certificate to the client.
	//
	//   - true: Enable SNI.
	//
	//   - false: Disable SNI.
	//
	//   followRedirect: Specifies whether to follow 3xx redirections.
	//
	//   - true: Follow the redirection if the detection point receives a 3xx status code (301, 302, 303, 307, or 308).
	//
	//   - false: Do not follow the redirection.
	//
	// - Ping:
	//
	//   packetNum: The number of ICMP packets to send for each ping health check. Valid values: 20, 50, and 100.
	//
	//   packetLossRate: The packet loss rate that triggers an alarm. For each ping health check, the system calculates the packet loss rate. Packet loss rate = (Number of lost packets / Total number of sent ICMP packets) × 100%. An alarm is triggered if the packet loss rate reaches this threshold. Valid values: 10, 30, 40, 80, 90, and 100.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The percentage of failed detection points. An endpoint is considered abnormal if the percentage of detection points that fail the health check exceeds this threshold. Valid values:
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
	// The interval between health checks in seconds. The default value is 60. The minimum interval is 15 seconds. This feature is available only for Ultimate Edition instances.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The list of detection points. Call [ListCloudGtmMonitorNodes](https://help.aliyun.com/document_detail/2797327.html) to obtain the information.
	IspCityNodes []*UpdateCloudGtmMonitorTemplateRequestIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Repeated"`
	// The name of the health check template. For easy identification, name the template based on its health check protocol.
	//
	// example:
	//
	// Ping-IPv4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique ID of the health check template that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// mtp-89518052425100****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The timeout period for a health check in milliseconds. If a packet is not returned within the specified timeout period, the health check fails. Valid values:
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

type UpdateCloudGtmMonitorTemplateRequestIspCityNodes struct {
	// The city code of the detection point.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The carrier code of the detection point.
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
