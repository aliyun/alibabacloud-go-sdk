// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackType(v int64) *DescribeEventLogRequest
	GetAttackType() *int64
	SetCurrentPage(v string) *DescribeEventLogRequest
	GetCurrentPage() *string
	SetDstIP(v string) *DescribeEventLogRequest
	GetDstIP() *string
	SetEndTime(v string) *DescribeEventLogRequest
	GetEndTime() *string
	SetFirewallType(v string) *DescribeEventLogRequest
	GetFirewallType() *string
	SetLang(v string) *DescribeEventLogRequest
	GetLang() *string
	SetLogType(v string) *DescribeEventLogRequest
	GetLogType() *string
	SetPageSize(v string) *DescribeEventLogRequest
	GetPageSize() *string
	SetRuleResult(v string) *DescribeEventLogRequest
	GetRuleResult() *string
	SetSourceCode(v string) *DescribeEventLogRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeEventLogRequest
	GetSourceIp() *string
	SetSrcIP(v string) *DescribeEventLogRequest
	GetSrcIP() *string
	SetStartTime(v string) *DescribeEventLogRequest
	GetStartTime() *string
}

type DescribeEventLogRequest struct {
	// The attack type of the intrusion prevention event. Valid values:
	//
	// - **1**: Abnormal connection.
	//
	// - **2**: Command execution.
	//
	// - **3**: Brute-force attack.
	//
	// - **4**: Scanning.
	//
	// - **5**: Other.
	//
	// - **6**: Information leakage.
	//
	// - **7**: DoS attack.
	//
	// - **8**: Overflow attack.
	//
	// - **9**: Web attack.
	//
	// - **10**: Trojan backdoor.
	//
	// - **11**: Viruses and worms.
	//
	// - **12**: Mining behavior.
	//
	// - **13**: Reverse shell.
	//
	// > If you do not specify this parameter, all attack types are queried.
	//
	// example:
	//
	// 13
	AttackType *int64 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The page number for a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The destination IP address to query. If you specify this parameter, the system queries intrusion prevention events that involve this destination IP address.
	//
	// example:
	//
	// 10.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1731551104
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the Cloud Firewall. Valid values:
	//
	// - **VpcFirewall**: virtual private cloud (VPC) firewalls.
	//
	// - **InternetFirewall*	- (default): the Internet Border firewall.
	//
	// example:
	//
	// [
	//
	//   "internet",
	//
	//   "vpc",
	//
	//   "nat"
	//
	// ]
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The language of the response message.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The log type.
	//
	// example:
	//
	// 1
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **6**, which indicates that 6 entries are returned on each page. Maximum value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The defense status of the Cloud Firewall. Valid values:
	//
	// - **1**: Alert.
	//
	// - **2**: Block.
	//
	// > If you do not specify this parameter, all defense statuses are queried.
	//
	// example:
	//
	// 0
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The tracing code.
	//
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// The source IP address of the requester.
	//
	// example:
	//
	// 60.182.79.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The source IP address to query. If you specify this parameter, the system queries intrusion prevention events that involve this source IP address.
	//
	// example:
	//
	// 47.92.x.x
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1692504764000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEventLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventLogRequest) GetAttackType() *int64 {
	return s.AttackType
}

func (s *DescribeEventLogRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeEventLogRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeEventLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEventLogRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribeEventLogRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventLogRequest) GetLogType() *string {
	return s.LogType
}

func (s *DescribeEventLogRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeEventLogRequest) GetRuleResult() *string {
	return s.RuleResult
}

func (s *DescribeEventLogRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeEventLogRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeEventLogRequest) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeEventLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEventLogRequest) SetAttackType(v int64) *DescribeEventLogRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeEventLogRequest) SetCurrentPage(v string) *DescribeEventLogRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventLogRequest) SetDstIP(v string) *DescribeEventLogRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeEventLogRequest) SetEndTime(v string) *DescribeEventLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventLogRequest) SetFirewallType(v string) *DescribeEventLogRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeEventLogRequest) SetLang(v string) *DescribeEventLogRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventLogRequest) SetLogType(v string) *DescribeEventLogRequest {
	s.LogType = &v
	return s
}

func (s *DescribeEventLogRequest) SetPageSize(v string) *DescribeEventLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventLogRequest) SetRuleResult(v string) *DescribeEventLogRequest {
	s.RuleResult = &v
	return s
}

func (s *DescribeEventLogRequest) SetSourceCode(v string) *DescribeEventLogRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeEventLogRequest) SetSourceIp(v string) *DescribeEventLogRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeEventLogRequest) SetSrcIP(v string) *DescribeEventLogRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeEventLogRequest) SetStartTime(v string) *DescribeEventLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventLogRequest) Validate() error {
	return dara.Validate(s)
}
