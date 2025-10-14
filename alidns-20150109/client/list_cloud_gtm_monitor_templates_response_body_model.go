// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmMonitorTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCloudGtmMonitorTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmMonitorTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCloudGtmMonitorTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v *ListCloudGtmMonitorTemplatesResponseBodyTemplates) *ListCloudGtmMonitorTemplatesResponseBody
	GetTemplates() *ListCloudGtmMonitorTemplatesResponseBodyTemplates
	SetTotalItems(v int32) *ListCloudGtmMonitorTemplatesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListCloudGtmMonitorTemplatesResponseBody
	GetTotalPages() *int32
}

type ListCloudGtmMonitorTemplatesResponseBody struct {
	// Current page number, starting from 1, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of 100 and a default of 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The health check templates.
	Templates *ListCloudGtmMonitorTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	// Total number of health check template entries retrieved.
	//
	// example:
	//
	// 30
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages after data pagination.
	//
	// example:
	//
	// 2
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCloudGtmMonitorTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetTemplates() *ListCloudGtmMonitorTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetPageNumber(v int32) *ListCloudGtmMonitorTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetPageSize(v int32) *ListCloudGtmMonitorTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetRequestId(v string) *ListCloudGtmMonitorTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetTemplates(v *ListCloudGtmMonitorTemplatesResponseBodyTemplates) *ListCloudGtmMonitorTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetTotalItems(v int32) *ListCloudGtmMonitorTemplatesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) SetTotalPages(v int32) *ListCloudGtmMonitorTemplatesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		if err := s.Templates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmMonitorTemplatesResponseBodyTemplates struct {
	Template []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplates) GetTemplate() []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	return s.Template
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplates) SetTemplate(v []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) *ListCloudGtmMonitorTemplatesResponseBodyTemplates {
	s.Template = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplates) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate struct {
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
	// The number of retries. The system will only judge the application service as abnormal after consecutive monitoring failures to prevent inaccurate monitoring results due to momentary network fluctuations or other reasons. Available retry options are:
	//
	// - 1 - 2 - 3
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
	// Percentage of selected node probe failures (%), that is, the percentage of unhealthy check points among total probe points. When the failure ratio exceeds the set threshold, the service address is judged as abnormal. The available failure ratio thresholds are:
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
	// The time interval between each check (in seconds), with a default of probing once every minute. The minimum supported health check interval is 15 seconds, available for flagship edition instances.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The IP address type of health check nodes. Valid values:
	//
	// 	- IPv4: applicable when health checks are performed on IPv4 addresses.
	//
	// 	- IPv6: applicable when health checks are performed on IPv6 addresses.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The health check nodes. You can call the [ListCloudGtmMonitorNodes](~~ListCloudGtmMonitorNodes~~) operation to obtain the health check nodes.
	IspCityNodes *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Struct"`
	// The name of the health check probe template, generally for the convenience of configuration personnel to distinguish and remember.
	//
	// example:
	//
	// IPv4-Ping
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protocol types for initiating probes to the target IP address:
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
	// Last modification time of the health check template.
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

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetFailureRate() *int32 {
	return s.FailureRate
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetInterval() *int32 {
	return s.Interval
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetIspCityNodes() *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes {
	return s.IspCityNodes
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetName() *string {
	return s.Name
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetProtocol() *string {
	return s.Protocol
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetRemark() *string {
	return s.Remark
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetCreateTime(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.CreateTime = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetCreateTimestamp(v int64) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetEvaluationCount(v int32) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.EvaluationCount = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetExtendInfo(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.ExtendInfo = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetFailureRate(v int32) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.FailureRate = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetInterval(v int32) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Interval = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetIpVersion(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.IpVersion = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetIspCityNodes(v *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.IspCityNodes = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Name = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetProtocol(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Protocol = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetRemark(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Remark = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetTemplateId(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.TemplateId = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetTimeout(v int32) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.Timeout = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetUpdateTime(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) SetUpdateTimestamp(v int64) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplate) Validate() error {
	if s.IspCityNodes != nil {
		if err := s.IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes struct {
	IspCityNode []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) GetIspCityNode() []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	return s.IspCityNode
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) SetIspCityNode(v []*ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes {
	s.IspCityNode = v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodes) Validate() error {
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

type ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode struct {
	// City code.
	//
	// example:
	//
	// 738
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// City name.
	//
	// example:
	//
	// Beijing
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// Country code.
	//
	// example:
	//
	// 629
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// Country name.
	//
	// example:
	//
	// China
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// Probe node group name.
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
	// Operator code.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// Operator name.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GoString() string {
	return s.String()
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCountryCode() *string {
	return s.CountryCode
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetCountryName() *string {
	return s.CountryName
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetGroupName() *string {
	return s.GroupName
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetGroupType() *string {
	return s.GroupType
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCityCode(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CityCode = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCityName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CityName = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCountryCode(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CountryCode = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetCountryName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.CountryName = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetGroupName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.GroupName = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetGroupType(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.GroupType = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetIspCode(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.IspCode = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) SetIspName(v string) *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode {
	s.IspName = &v
	return s
}

func (s *ListCloudGtmMonitorTemplatesResponseBodyTemplatesTemplateIspCityNodesIspCityNode) Validate() error {
	return dara.Validate(s)
}
