// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmMonitorTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeCloudGtmMonitorTemplateResponseBody
	GetCreateTimestamp() *int64
	SetEvaluationCount(v int32) *DescribeCloudGtmMonitorTemplateResponseBody
	GetEvaluationCount() *int32
	SetExtendInfo(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetExtendInfo() *string
	SetFailureRate(v int32) *DescribeCloudGtmMonitorTemplateResponseBody
	GetFailureRate() *int32
	SetInterval(v int32) *DescribeCloudGtmMonitorTemplateResponseBody
	GetInterval() *int32
	SetIpVersion(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetIpVersion() *string
	SetIspCityNodes(v *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes) *DescribeCloudGtmMonitorTemplateResponseBody
	GetIspCityNodes() *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes
	SetName(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetName() *string
	SetProtocol(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetProtocol() *string
	SetRemark(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetTemplateId() *string
	SetTimeout(v int32) *DescribeCloudGtmMonitorTemplateResponseBody
	GetTimeout() *int32
	SetUpdateTime(v string) *DescribeCloudGtmMonitorTemplateResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeCloudGtmMonitorTemplateResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeCloudGtmMonitorTemplateResponseBody struct {
	// Health check template creation time.
	//
	// example:
	//
	// 2024-03-23T13:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Health check template creation time (timestamp).
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Retries count. The system will only judge the application service as abnormal after consecutive monitoring failures to prevent inaccurate monitoring results due to momentary network fluctuations or other reasons. Available retry counts are:
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
	//     **host**: the Host field of an HTTP or HTTPS request header during an HTTP or HTTPS health check. The parameter value indicates the HTTP website that you want to visit. By default, the value is the primary domain name. You can change the value based on your business requirements.
	//
	//     **path**: the URL for HTTP or HTTPS health checks. Default value: /.
	//
	//     **code**: the alert threshold. During an HTTP or HTTPS health check, the system checks whether a web server functions as expected based on the status code that is returned from the web server. If the returned status code is greater than the specified threshold, the corresponding application service address is deemed abnormal. Valid values:
	//
	//     	- 400: indicates an invalid request. If an HTTP or HTTPS request contains invalid request parameters, a web server returns a status code that is greater than 400. You must specify an exact URL for path if you set code to 400.
	//
	//     	- 500: indicates a server error. If some exceptions occur on a web server, the web server returns a status code that is greater than 500. This value is used by default.
	//
	//     **sni**: indicates whether Server Name Indication (SNI) is enabled. This parameter is used only when the health check protocol is HTTPS. SNI is an extension to the Transport Layer Security (TLS) protocol, which allows a client to specify the host to be connected when the client sends a TLS handshake request. TLS handshakes occur before any data of HTTP requests is sent. Therefore, SNI enables servers to identify the services that clients are attempting to access before certificates are sent. This allows the servers to present correct certificates to the clients. Valid values:
	//
	//     	- true: SNI is enabled.
	//
	//     	- false: SNI is disabled.
	//
	//     **followRedirect**: indicates whether 3XX redirects are followed. Valid values:
	//
	//     	- true: 3XX redirects are followed. You are redirected to the destination address if a 3XX status code such as 301, 302, 303, 307, or 308 is returned.
	//
	//     	- false: 3XX redirects are not followed.
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
	// The time interval (in seconds) between each check, with a default interval of 1 minute. The minimum supported health check interval is 15 seconds, available for flagship edition instances.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Detect the type of the node IP address:
	//
	// - IPv4: Applicable when the target address type is IPv4;
	//
	// - IPv6: Applicable when the target address type is IPv6.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Probe node list, detailed information can be obtained by calling ListCloudGtmMonitorNodes.
	IspCityNodes *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Struct"`
	// The name of the health check probe template, which is recommended to be distinguishable for configuration personnel to differentiate and remember, ideally indicating the health check protocol.
	//
	// example:
	//
	// Ping-IPv4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protocol types to initiate probes to the target IP address:
	//
	// - ping
	//
	// - tcp
	//
	// - http
	//
	// - https
	//
	// example:
	//
	// ping
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Remarks for the health check template.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 6AEC7A64-3CB1-4C49-8B35-0B901F1E26BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the health check template. This ID uniquely identifies the health check template.
	//
	// example:
	//
	// mtp-89518052425100**80
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Probe timeout (in milliseconds), data packets not returned within the timeout period are deemed as health check timeouts:
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
	// Health check template configuration modification time.
	//
	// example:
	//
	// 2024-03-29T13:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Health check template configuration modification time (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeCloudGtmMonitorTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmMonitorTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetFailureRate() *int32 {
	return s.FailureRate
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetIspCityNodes() *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes {
	return s.IspCityNodes
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetCreateTime(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetCreateTimestamp(v int64) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetEvaluationCount(v int32) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetExtendInfo(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.ExtendInfo = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetFailureRate(v int32) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.FailureRate = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetInterval(v int32) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetIpVersion(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.IpVersion = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetIspCityNodes(v *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.IspCityNodes = v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetName(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetProtocol(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetRemark(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetRequestId(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetTemplateId(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetTimeout(v int32) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetUpdateTime(v string) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) SetUpdateTimestamp(v int64) *DescribeCloudGtmMonitorTemplateResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes struct {
	IspCityNode []*DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes) GetIspCityNode() []*DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	return s.IspCityNode
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes) SetIspCityNode(v []*DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes {
	s.IspCityNode = v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode struct {
	// City code
	//
	// example:
	//
	// 357
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// City name
	//
	// example:
	//
	// Shanghai
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Country Code
	//
	// example:
	//
	// 629
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// Country Name
	//
	// example:
	//
	// China
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// Probe node group type name
	//
	// example:
	//
	// BGP Nodes
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Probe node group types:
	//
	// - BGP: BGP nodes
	//
	// - OVERSEAS: International nodes
	//
	// - ISP: Carrier nodes
	//
	// example:
	//
	// BGP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// Operator Code
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// Operator Name
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GetCountryName() *string {
	return s.CountryName
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) SetCityCode(v string) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	s.CityCode = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) SetCityName(v string) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	s.CityName = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) SetCountryCode(v string) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	s.CountryCode = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) SetCountryName(v string) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	s.CountryName = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) SetGroupName(v string) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	s.GroupName = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) SetGroupType(v string) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	s.GroupType = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) SetIspCode(v string) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	s.IspCode = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) SetIspName(v string) *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode {
	s.IspName = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode) Validate() error {
	return dara.Validate(s)
}
