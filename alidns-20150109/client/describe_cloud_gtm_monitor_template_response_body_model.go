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
	// The time when the health check template was created.
	//
	// example:
	//
	// 2024-03-23T13:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The UNIX timestamp that indicates when the health check template was created.
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The number of consecutive times that a health check must fail before the application service is declared abnormal. This prevents false alarms caused by transient issues such as network jitter. Valid values:
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
	// The extended information in a JSON string. The parameters vary based on the protocol.
	//
	// - For HTTP and HTTPS:
	//
	//   **host**: The Host field in the HTTP or HTTPS request header. This field identifies the website that you want to access. The default value is the primary domain name. If the target website has specific host requirements, modify this parameter.
	//
	//   **path**: The URL path for HTTP or HTTPS health checks. The default value is /.
	//
	//   **code**: The system determines whether the web server is working as expected based on the return code. If the return code is greater than the specified threshold, the application service is considered abnormal.
	//
	//   - 400: Bad Request. If an HTTP or HTTPS request contains incorrect parameters, the web server returns a code greater than 400. If you set the threshold to 400, specify the exact URL path.
	//
	//   - 500: Server Error. If an exception occurs on the web server, it returns a code greater than 500. The default threshold is 500.
	//
	//   **sni**: Specifies whether to enable Server Name Indication (SNI). This parameter is used only for the HTTPS protocol. SNI is a Transport Layer Security (TLS) extension that allows a client to specify the hostname it wants to connect to at the start of the TLS handshake. This allows the server to present the correct certificate for that hostname.
	//
	//   - true: Enable SNI.
	//
	//   - false: Disable SNI.
	//
	//   **followRedirect**: Specifies whether to follow 3xx redirections.
	//
	//   - true: If the status code returned by the detection point is 3xx (301, 302, 303, 307, or 308), the system follows the redirection.
	//
	//   - false: The system does not follow the redirection.
	//
	// - For ping:
	//
	//   **packetNum**: The number of ICMP packets to send for each ping health check. Valid values: 20, 50, and 100.
	//
	//   **packetLossRate**: The packet loss rate threshold. For each ping health check, the system calculates the packet loss rate. If the packet loss rate reaches the threshold, an alert is triggered. Packet loss rate = (Number of lost packets / Total number of sent ICMP packets) × 100%. Valid values for the packet loss rate are 10, 30, 40, 80, 90, and 100.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The percentage of failed detection points. If the percentage of failed detection points exceeds this value, the endpoint is declared abnormal. Valid values:
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
	// The IP address type of the detection points:
	//
	// - IPv4: The target address is an IPv4 address.
	//
	// - IPv6: The target address is an IPv6 address.
	//
	// example:
	//
	// IPv4
	IpVersion    *string                                                  `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IspCityNodes *DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Struct"`
	// The name of the health check template. To easily identify the template, specify a name that indicates the health check protocol.
	//
	// example:
	//
	// Ping-IPv4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protocol used to probe the target IP address:
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
	// The remarks on the health check template.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 6AEC7A64-3CB1-4C49-8B35-0B901F1E26BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unique ID of the health check template.
	//
	// example:
	//
	// mtp-89518052425100****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The health check timeout period in milliseconds. If a packet is not returned within the specified timeout period, the health check fails. Valid values:
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
	// The time when the health check template was last modified.
	//
	// example:
	//
	// 2024-03-29T13:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The UNIX timestamp that indicates when the health check template was last modified.
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
	if s.IspCityNodes != nil {
		if err := s.IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.IspCityNode != nil {
		for _, item := range s.IspCityNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudGtmMonitorTemplateResponseBodyIspCityNodesIspCityNode struct {
	CityCode    *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	CityName    *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType   *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IspCode     *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	IspName     *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
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
