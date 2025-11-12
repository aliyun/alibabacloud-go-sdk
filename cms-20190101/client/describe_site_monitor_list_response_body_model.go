// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSiteMonitorListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeSiteMonitorListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeSiteMonitorListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSiteMonitorListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSiteMonitorListResponseBody
	GetRequestId() *string
	SetSiteMonitors(v *DescribeSiteMonitorListResponseBodySiteMonitors) *DescribeSiteMonitorListResponseBody
	GetSiteMonitors() *DescribeSiteMonitorListResponseBodySiteMonitors
	SetSuccess(v string) *DescribeSiteMonitorListResponseBody
	GetSuccess() *string
	SetTotalCount(v int32) *DescribeSiteMonitorListResponseBody
	GetTotalCount() *int32
}

type DescribeSiteMonitorListResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A80DB41C-AF6C-50E1-ADB5-66DCBA3D266B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The site monitoring tasks that are returned.
	SiteMonitors *DescribeSiteMonitorListResponseBodySiteMonitors `json:"SiteMonitors,omitempty" xml:"SiteMonitors,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSiteMonitorListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSiteMonitorListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSiteMonitorListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSiteMonitorListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSiteMonitorListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteMonitorListResponseBody) GetSiteMonitors() *DescribeSiteMonitorListResponseBodySiteMonitors {
	return s.SiteMonitors
}

func (s *DescribeSiteMonitorListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSiteMonitorListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSiteMonitorListResponseBody) SetCode(v string) *DescribeSiteMonitorListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetMessage(v string) *DescribeSiteMonitorListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetPageNumber(v int32) *DescribeSiteMonitorListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetPageSize(v int32) *DescribeSiteMonitorListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetRequestId(v string) *DescribeSiteMonitorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetSiteMonitors(v *DescribeSiteMonitorListResponseBodySiteMonitors) *DescribeSiteMonitorListResponseBody {
	s.SiteMonitors = v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetSuccess(v string) *DescribeSiteMonitorListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetTotalCount(v int32) *DescribeSiteMonitorListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) Validate() error {
	if s.SiteMonitors != nil {
		if err := s.SiteMonitors.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorListResponseBodySiteMonitors struct {
	SiteMonitor []*DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor `json:"SiteMonitor,omitempty" xml:"SiteMonitor,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorListResponseBodySiteMonitors) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBodySiteMonitors) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitors) GetSiteMonitor() []*DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	return s.SiteMonitor
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitors) SetSiteMonitor(v []*DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) *DescribeSiteMonitorListResponseBodySiteMonitors {
	s.SiteMonitor = v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitors) Validate() error {
	if s.SiteMonitor != nil {
		for _, item := range s.SiteMonitor {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor struct {
	// The URL or IP address that is monitored by the site monitoring task.
	//
	// example:
	//
	// https://aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The detection point type. Valid values:
	//
	// 	- PC
	//
	// 	- MOBILE
	//
	// example:
	//
	// PC
	AgentGroup *string `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	// The time when the site monitoring task was created.
	//
	// example:
	//
	// 2021-11-01 11:05:18
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The interval at which detection requests are sent. Unit: minutes.
	//
	// example:
	//
	// 1
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The extended options of the site monitoring task. The options vary based on the specified protocol. For more information, see [CreateSiteMonitor](https://help.aliyun.com/document_detail/115048.html).
	OptionsJson *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty" type:"Struct"`
	// The ID of the site monitoring task.
	//
	// example:
	//
	// f5783760-1b39-4b6b-80e8-453d962a****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the site monitoring task.
	//
	// example:
	//
	// 域名检测_example.com.cn
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task status. Valid values:
	//
	// 	- 1: The task is enabled.
	//
	// 	- 2: The task is disabled.
	//
	// example:
	//
	// 1
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The protocol that is used by the site monitoring task. Valid values: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	//
	// example:
	//
	// TCP
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The time when the site monitoring task was updated.
	//
	// example:
	//
	// 2022-03-08 17:14:31
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetAddress() *string {
	return s.Address
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetAgentGroup() *string {
	return s.AgentGroup
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetOptionsJson() *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	return s.OptionsJson
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetAddress(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.Address = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetAgentGroup(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.AgentGroup = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetCreateTime(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.CreateTime = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetInterval(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.Interval = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetOptionsJson(v *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.OptionsJson = v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetTaskId(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetTaskName(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.TaskName = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetTaskState(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.TaskState = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetTaskType(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.TaskType = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetUpdateTime(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) Validate() error {
	if s.OptionsJson != nil {
		if err := s.OptionsJson.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson struct {
	// The acceptable status code.
	//
	// >  We recommend that you configure assertions.
	//
	// example:
	//
	// 400
	AcceptableResponseCode *string `json:"acceptable_response_code,omitempty" xml:"acceptable_response_code,omitempty"`
	// The assertions.
	Assertions *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions `json:"assertions,omitempty" xml:"assertions,omitempty" type:"Struct"`
	// The number of retries after a DNS failure occurred.
	//
	// example:
	//
	// 3
	Attempts *int64 `json:"attempts,omitempty" xml:"attempts,omitempty"`
	// Indicates whether the security authentication feature is enabled. Valid values:
	//
	// 	- 0: The feature is enabled.
	//
	// 	- 1: The feature is disabled.
	//
	// example:
	//
	// 1
	Authentication *int32 `json:"authentication,omitempty" xml:"authentication,omitempty"`
	// Indicates whether the certificate is verified. Valid values:
	//
	// 	- false (default): The certificate is not verified.
	//
	// 	- true: The certificate is verified.
	//
	// example:
	//
	// false
	CertVerify *bool `json:"cert_verify,omitempty" xml:"cert_verify,omitempty"`
	// The cookie of the HTTP request.
	//
	// example:
	//
	// lang=en
	Cookie *string `json:"cookie,omitempty" xml:"cookie,omitempty"`
	// Indicates whether MTR is automatically used to diagnose network issues if a task fails. Valid values:
	//
	// 	- false (default): MTR is not automatically used to diagnose network issues if a task fails.
	//
	// 	- true: MTR is automatically used to diagnose network issues if a task fails.
	//
	// example:
	//
	// false
	DiagnosisMtr *bool `json:"diagnosis_mtr,omitempty" xml:"diagnosis_mtr,omitempty"`
	// Indicates whether ping requests are automatically sent to detect network latency if a detection task fails. Valid values:
	//
	// 	- false (default): Ping requests are not automatically sent to detect network latency if a detection task fails.
	//
	// 	- true: Ping requests are automatically sent to detect network latency if a detection task fails.
	//
	// example:
	//
	// false
	DiagnosisPing *bool `json:"diagnosis_ping,omitempty" xml:"diagnosis_ping,omitempty"`
	// The relationship between the list of expected aliases or IP addresses and the list of DNS results. Valid values:
	//
	// 	- IN_DNS: The list of expected values is a subset of the list of DNS results.
	//
	// 	- DNS_IN: The list of DNS results is a subset of the list of expected values.
	//
	// 	- EQUAL: The list of DNS results is the same as the list of expected values.
	//
	// 	- ANY: The list of DNS results intersects with the list of expected values.
	//
	// example:
	//
	// IN_DNS
	DnsMatchRule *string `json:"dns_match_rule,omitempty" xml:"dns_match_rule,omitempty"`
	// The domain name or IP address of the DNS server.
	//
	// example:
	//
	// 192.168.XX.XX
	DnsServer *string `json:"dns_server,omitempty" xml:"dns_server,omitempty"`
	// The type of the DNS record. This parameter is returned only if the TaskType parameter is set to DNS. Valid values:
	//
	// 	- A (default): a record that specifies an IP address related to the specified host name or domain name.
	//
	// 	- CNAME: a record that maps multiple domain names to a domain name.
	//
	// 	- NS: a record that specifies a DNS server used to parse domain names.
	//
	// 	- MX: a record that links domain names to the address of a mail server.
	//
	// 	- TXT: a record that stores the text information of host name or domain names. The text must be 1 to 512 bytes in length. The TXT record serves as a Sender Policy Framework (SPF) record to fight against spam.
	//
	// 	- AAAA: a record that maps a domain name to the relevant IPv6 address.
	//
	// example:
	//
	// A
	DnsType *string `json:"dns_type,omitempty" xml:"dns_type,omitempty"`
	// Indicates whether the DNS server of the carrier is used.
	//
	// 	- true (default): The DNS server of the carrier is used.
	//
	// 	- false: The DNS server of the carrier is not used. The default DNS server or the specified DNS server is used.
	//
	// example:
	//
	// true
	EnableOperatorDns *bool `json:"enable_operator_dns,omitempty" xml:"enable_operator_dns,omitempty"`
	// The packet loss rate.
	//
	// >  This parameter is returned only if the TaskType parameter is set to PING.
	//
	// example:
	//
	// 0.5
	FailureRate *float32 `json:"failure_rate,omitempty" xml:"failure_rate,omitempty"`
	// The header of the HTTP request. An HTTP header is a key-value pair in which the key and the value are separated by a colon (:). The format is `key1:value1`. Each HTTP header occupies a line.
	//
	// example:
	//
	// testKey:testValue
	Header *string `json:"header,omitempty" xml:"header,omitempty"`
	// The HTTP request method. Valid values:
	//
	// 	- get
	//
	// 	- post
	//
	// 	- head
	//
	// example:
	//
	// get
	HttpMethod *string `json:"http_method,omitempty" xml:"http_method,omitempty"`
	// Indicates whether the password is decoded by using the Base64 algorithm. Valid values:
	//
	// 	- true: The password is decoded by using the Base64 algorithm.
	//
	// 	- false (default): The password is not decoded by using the Base64 algorithm.
	//
	// example:
	//
	// false
	IsBase64Encode *string `json:"isBase64Encode,omitempty" xml:"isBase64Encode,omitempty"`
	// Indicates whether the alert rule is included. Valid values:
	//
	// 	- 0: The alert rule is included.
	//
	// 	- 1: The alert rule is not included.
	//
	// example:
	//
	// 0
	MatchRule *int32 `json:"match_rule,omitempty" xml:"match_rule,omitempty"`
	// The password of the SMTP, POP3, or FTP protocol.
	//
	// example:
	//
	// 123****
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The number of hops for the PING protocol.
	//
	// example:
	//
	// 20
	PingNum *int32 `json:"ping_num,omitempty" xml:"ping_num,omitempty"`
	// The port number of the TCP, UDP, SMTP, or POP3 protocol.
	//
	// example:
	//
	// 80
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol type of DNS detection. Valid values:
	//
	// 	- udp (default)
	//
	// 	- tcp
	//
	// 	- tcp-tls
	//
	// example:
	//
	// udp
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Indicates whether the PROXY protocol is enabled. Valid values:
	//
	// 	- false (default): The PROXY protocol is disabled.
	//
	// 	- true: The PROXY protocol is enabled.
	//
	// example:
	//
	// false
	ProxyProtocol *bool `json:"proxy_protocol,omitempty" xml:"proxy_protocol,omitempty"`
	// The content of the HTTP request.
	//
	// example:
	//
	// cf0f85
	RequestContent *string `json:"request_content,omitempty" xml:"request_content,omitempty"`
	// The format of the HTTP request. Valid values:
	//
	// 	- hex: hexadecimal
	//
	// 	- txt: text
	//
	// example:
	//
	// hex
	RequestFormat *string `json:"request_format,omitempty" xml:"request_format,omitempty"`
	// The response to the HTTP request.
	//
	// 	- Hexadecimal format: If the request content is a byte string and cannot be represented in printable characters, you can convert the byte string to printable characters in the hexadecimal format. If you convert the byte string to printable characters in the hexadecimal format, one byte is converted to two hexadecimal characters. For example, (byte)1 is converted to `01` and (byte)27 is converted to `1B`. If the request content is a binary array in the Java format, for example, `{(byte)1, (byte)27}`, you can convert the binary array to `011b` or `011B`. Hexadecimal characters are not case-sensitive in site monitoring tasks. You can enter `011B` in the request content and set the request_format parameter to hex.
	//
	// 	- Text format: Common text refers to strings that consist of printable characters.
	//
	// example:
	//
	// cf0f85
	ResponseContent *string `json:"response_content,omitempty" xml:"response_content,omitempty"`
	// The format of the HTTP response. Valid values:
	//
	// 	- hex: hexadecimal
	//
	// 	- txt: text
	//
	// example:
	//
	// hex
	ResponseFormat *string `json:"response_format,omitempty" xml:"response_format,omitempty"`
	// The number of times a failed detection request is retried.
	//
	// example:
	//
	// 1
	RetryDelay *int32 `json:"retry_delay,omitempty" xml:"retry_delay,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	TimeOut *int64 `json:"time_out,omitempty" xml:"time_out,omitempty"`
	// Indicates whether redirects are followed if the status code 301 or 302 is returned. Valid values:
	//
	// 	- true: Redirects are not followed.
	//
	// 	- false (default): Redirects are followed.
	//
	// example:
	//
	// false
	UnfollowRedirect *bool `json:"unfollow_redirect,omitempty" xml:"unfollow_redirect,omitempty"`
	// The username of the FTP, SMTP, or POP3 protocol.
	//
	// example:
	//
	// testUser
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetAcceptableResponseCode() *string {
	return s.AcceptableResponseCode
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetAssertions() *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions {
	return s.Assertions
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetAttempts() *int64 {
	return s.Attempts
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetAuthentication() *int32 {
	return s.Authentication
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetCertVerify() *bool {
	return s.CertVerify
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetDiagnosisMtr() *bool {
	return s.DiagnosisMtr
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetDiagnosisPing() *bool {
	return s.DiagnosisPing
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetDnsMatchRule() *string {
	return s.DnsMatchRule
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetDnsServer() *string {
	return s.DnsServer
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetDnsType() *string {
	return s.DnsType
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetEnableOperatorDns() *bool {
	return s.EnableOperatorDns
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetFailureRate() *float32 {
	return s.FailureRate
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetHeader() *string {
	return s.Header
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetIsBase64Encode() *string {
	return s.IsBase64Encode
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetMatchRule() *int32 {
	return s.MatchRule
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetPassword() *string {
	return s.Password
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetPingNum() *int32 {
	return s.PingNum
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetPort() *int32 {
	return s.Port
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetProxyProtocol() *bool {
	return s.ProxyProtocol
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetRequestContent() *string {
	return s.RequestContent
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetRequestFormat() *string {
	return s.RequestFormat
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetResponseContent() *string {
	return s.ResponseContent
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetResponseFormat() *string {
	return s.ResponseFormat
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetRetryDelay() *int32 {
	return s.RetryDelay
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetTimeOut() *int64 {
	return s.TimeOut
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetUnfollowRedirect() *bool {
	return s.UnfollowRedirect
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GetUsername() *string {
	return s.Username
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetAcceptableResponseCode(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.AcceptableResponseCode = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetAssertions(v *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Assertions = v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetAttempts(v int64) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Attempts = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetAuthentication(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Authentication = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetCertVerify(v bool) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.CertVerify = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetCookie(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Cookie = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetDiagnosisMtr(v bool) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.DiagnosisMtr = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetDiagnosisPing(v bool) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.DiagnosisPing = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetDnsMatchRule(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.DnsMatchRule = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetDnsServer(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.DnsServer = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetDnsType(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.DnsType = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetEnableOperatorDns(v bool) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.EnableOperatorDns = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetFailureRate(v float32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.FailureRate = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetHeader(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Header = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetHttpMethod(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.HttpMethod = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetIsBase64Encode(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.IsBase64Encode = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetMatchRule(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.MatchRule = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetPassword(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Password = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetPingNum(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.PingNum = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetPort(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Port = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetProtocol(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Protocol = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetProxyProtocol(v bool) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.ProxyProtocol = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetRequestContent(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.RequestContent = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetRequestFormat(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.RequestFormat = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetResponseContent(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.ResponseContent = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetResponseFormat(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.ResponseFormat = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetRetryDelay(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.RetryDelay = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetTimeOut(v int64) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.TimeOut = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetUnfollowRedirect(v bool) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.UnfollowRedirect = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetUsername(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Username = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) Validate() error {
	if s.Assertions != nil {
		if err := s.Assertions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions struct {
	Assertions []*DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions `json:"assertions,omitempty" xml:"assertions,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions) GetAssertions() []*DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions {
	return s.Assertions
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions) SetAssertions(v []*DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions {
	s.Assertions = v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions) Validate() error {
	if s.Assertions != nil {
		for _, item := range s.Assertions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions struct {
	// The comparison operator of the assertion. Valid values:
	//
	// 	- contains: contains
	//
	// 	- doesNotContain: does not contain
	//
	// 	- matches: matches regular expressions
	//
	// 	- doesNotMatch: does not match regular expressions
	//
	// 	- is: equal to a numeric value or matches a character
	//
	// 	- isNot: not equal to
	//
	// 	- lessThan: less than
	//
	// 	- moreThan: greater than
	//
	// example:
	//
	// lessThan
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The parsing path of the assertion.
	//
	// 	- If the assertion type is `body_json`, the path is `json path`.
	//
	// 	- If the assertion type is `body_xml`, the path is `xml path`.
	//
	// example:
	//
	// json path
	Property *string `json:"property,omitempty" xml:"property,omitempty"`
	// The numeric value or character used for matching.
	//
	// example:
	//
	// 1000
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// The assertion type. Valid values:
	//
	// 	- response_time: checks whether the response time meets expectations.
	//
	// 	- status_code: checks whether the HTTP status code meets expectations.
	//
	// 	- header: checks whether the fields in the response header meet expectations.
	//
	// 	- body_text: checks whether the content in the response body meets expectations by using text matching.
	//
	// 	- body_json: checks whether the content in the response body meets expectations by using JSON parsing (JSONPath).
	//
	// 	- body_xml: checks whether the content in the response body meets expectations by using XML parsing (XPath).
	//
	// example:
	//
	// response_time
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) GetOperator() *string {
	return s.Operator
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) GetProperty() *string {
	return s.Property
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) GetTarget() *string {
	return s.Target
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) GetType() *string {
	return s.Type
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) SetOperator(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions {
	s.Operator = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) SetProperty(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions {
	s.Property = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) SetTarget(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions {
	s.Target = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) SetType(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions {
	s.Type = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertionsAssertions) Validate() error {
	return dara.Validate(s)
}
