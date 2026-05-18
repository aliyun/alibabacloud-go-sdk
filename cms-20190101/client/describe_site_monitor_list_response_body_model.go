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
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Address     *string                                                                `json:"Address,omitempty" xml:"Address,omitempty"`
	AgentGroup  *string                                                                `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	CreateTime  *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Interval    *string                                                                `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OptionsJson *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty" type:"Struct"`
	TaskId      *string                                                                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName    *string                                                                `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskState   *string                                                                `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	TaskType    *string                                                                `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UpdateTime  *string                                                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	AcceptableResponseCode *string                                                                          `json:"acceptable_response_code,omitempty" xml:"acceptable_response_code,omitempty"`
	Assertions             *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJsonAssertions `json:"assertions,omitempty" xml:"assertions,omitempty" type:"Struct"`
	Attempts               *int64                                                                           `json:"attempts,omitempty" xml:"attempts,omitempty"`
	Authentication         *int32                                                                           `json:"authentication,omitempty" xml:"authentication,omitempty"`
	CertVerify             *bool                                                                            `json:"cert_verify,omitempty" xml:"cert_verify,omitempty"`
	Cookie                 *string                                                                          `json:"cookie,omitempty" xml:"cookie,omitempty"`
	DiagnosisMtr           *bool                                                                            `json:"diagnosis_mtr,omitempty" xml:"diagnosis_mtr,omitempty"`
	DiagnosisPing          *bool                                                                            `json:"diagnosis_ping,omitempty" xml:"diagnosis_ping,omitempty"`
	DnsMatchRule           *string                                                                          `json:"dns_match_rule,omitempty" xml:"dns_match_rule,omitempty"`
	DnsServer              *string                                                                          `json:"dns_server,omitempty" xml:"dns_server,omitempty"`
	DnsType                *string                                                                          `json:"dns_type,omitempty" xml:"dns_type,omitempty"`
	EnableOperatorDns      *bool                                                                            `json:"enable_operator_dns,omitempty" xml:"enable_operator_dns,omitempty"`
	FailureRate            *float32                                                                         `json:"failure_rate,omitempty" xml:"failure_rate,omitempty"`
	Header                 *string                                                                          `json:"header,omitempty" xml:"header,omitempty"`
	HttpMethod             *string                                                                          `json:"http_method,omitempty" xml:"http_method,omitempty"`
	IsBase64Encode         *string                                                                          `json:"isBase64Encode,omitempty" xml:"isBase64Encode,omitempty"`
	MatchRule              *int32                                                                           `json:"match_rule,omitempty" xml:"match_rule,omitempty"`
	Password               *string                                                                          `json:"password,omitempty" xml:"password,omitempty"`
	PingNum                *int32                                                                           `json:"ping_num,omitempty" xml:"ping_num,omitempty"`
	Port                   *int32                                                                           `json:"port,omitempty" xml:"port,omitempty"`
	Protocol               *string                                                                          `json:"protocol,omitempty" xml:"protocol,omitempty"`
	ProxyProtocol          *bool                                                                            `json:"proxy_protocol,omitempty" xml:"proxy_protocol,omitempty"`
	RequestContent         *string                                                                          `json:"request_content,omitempty" xml:"request_content,omitempty"`
	RequestFormat          *string                                                                          `json:"request_format,omitempty" xml:"request_format,omitempty"`
	ResponseContent        *string                                                                          `json:"response_content,omitempty" xml:"response_content,omitempty"`
	ResponseFormat         *string                                                                          `json:"response_format,omitempty" xml:"response_format,omitempty"`
	RetryDelay             *int32                                                                           `json:"retry_delay,omitempty" xml:"retry_delay,omitempty"`
	TimeOut                *int64                                                                           `json:"time_out,omitempty" xml:"time_out,omitempty"`
	UnfollowRedirect       *bool                                                                            `json:"unfollow_redirect,omitempty" xml:"unfollow_redirect,omitempty"`
	Username               *string                                                                          `json:"username,omitempty" xml:"username,omitempty"`
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
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Property *string `json:"property,omitempty" xml:"property,omitempty"`
	Target   *string `json:"target,omitempty" xml:"target,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
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
