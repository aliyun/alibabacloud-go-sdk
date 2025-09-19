// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAntiBruteForceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeAntiBruteForceRulesResponseBodyPageInfo) *DescribeAntiBruteForceRulesResponseBody
	GetPageInfo() *DescribeAntiBruteForceRulesResponseBodyPageInfo
	SetRequestId(v string) *DescribeAntiBruteForceRulesResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeAntiBruteForceRulesResponseBodyRules) *DescribeAntiBruteForceRulesResponseBody
	GetRules() []*DescribeAntiBruteForceRulesResponseBodyRules
}

type DescribeAntiBruteForceRulesResponseBody struct {
	// The pagination information.
	PageInfo *DescribeAntiBruteForceRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4E5BFDCF-B9DD-430D-9DA4-151BCB581C9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the defense rules returned.
	Rules []*DescribeAntiBruteForceRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeAntiBruteForceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponseBody) GetPageInfo() *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeAntiBruteForceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAntiBruteForceRulesResponseBody) GetRules() []*DescribeAntiBruteForceRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeAntiBruteForceRulesResponseBody) SetPageInfo(v *DescribeAntiBruteForceRulesResponseBodyPageInfo) *DescribeAntiBruteForceRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBody) SetRequestId(v string) *DescribeAntiBruteForceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBody) SetRules(v []*DescribeAntiBruteForceRulesResponseBodyRules) *DescribeAntiBruteForceRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAntiBruteForceRulesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAntiBruteForceRulesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) SetCount(v int32) *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) SetPageSize(v int32) *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeAntiBruteForceRulesResponseBodyRules struct {
	// 防暴力破解规则创建时间戳。单位：毫秒。
	//
	// example:
	//
	// 1669800181000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Indicates whether the defense rule is the default rule. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  The default rule takes effect on all servers that are not protected by defense rules against brute-force attacks.
	//
	// example:
	//
	// true
	DefaultRule *bool `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// false
	EnableSmartRule *bool `json:"EnableSmartRule,omitempty" xml:"EnableSmartRule,omitempty"`
	// The threshold of logon failures that you specify.
	//
	// example:
	//
	// 15
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The period of time during which logons from an account are not allowed. Unit: minutes.
	//
	// example:
	//
	// 360
	ForbiddenTime *int32 `json:"ForbiddenTime,omitempty" xml:"ForbiddenTime,omitempty"`
	// The ID of the defense rule.
	//
	// example:
	//
	// 1629
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of servers to which the defense rule is applied.
	//
	// example:
	//
	// 3
	MachineCount *int32 `json:"MachineCount,omitempty" xml:"MachineCount,omitempty"`
	// The name of the defense rule.
	//
	// example:
	//
	// AntiBruteForceRule01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The types of protocols that the brute force cracking rule supports to intercept.
	ProtocolType *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty" type:"Struct"`
	// The period of time during which logon failures from an account are measured. Unit: minutes. If **Span*	- is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
	//
	// example:
	//
	// 10
	Span *int32 `json:"Span,omitempty" xml:"Span,omitempty"`
	// An array consisting of the UUIDs of servers to which the defense rule is applied.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s DescribeAntiBruteForceRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetDefaultRule() *bool {
	return s.DefaultRule
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetEnableSmartRule() *bool {
	return s.EnableSmartRule
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetFailCount() *int32 {
	return s.FailCount
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetForbiddenTime() *int32 {
	return s.ForbiddenTime
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetId() *int64 {
	return s.Id
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetMachineCount() *int32 {
	return s.MachineCount
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetName() *string {
	return s.Name
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetProtocolType() *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType {
	return s.ProtocolType
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetSpan() *int32 {
	return s.Span
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) GetUuidList() []*string {
	return s.UuidList
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetCreateTimestamp(v int64) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetDefaultRule(v bool) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.DefaultRule = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetEnableSmartRule(v bool) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.EnableSmartRule = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetFailCount(v int32) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.FailCount = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetForbiddenTime(v int32) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.ForbiddenTime = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetId(v int64) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.Id = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetMachineCount(v int32) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.MachineCount = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetName(v string) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetProtocolType(v *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.ProtocolType = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetSpan(v int32) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.Span = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetUuidList(v []*string) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.UuidList = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}

type DescribeAntiBruteForceRulesResponseBodyRulesProtocolType struct {
	// RDP interception method, values:
	//
	// - **on**: enable
	//
	// - **off**: disable
	//
	// example:
	//
	// on
	Rdp *string `json:"Rdp,omitempty" xml:"Rdp,omitempty"`
	// SqlServer interception mode, with values:
	//
	// - **on**: enable
	//
	// - **off**: disable
	//
	// example:
	//
	// off
	SqlServer *string `json:"SqlServer,omitempty" xml:"SqlServer,omitempty"`
	// SSH interception method, with values:
	//
	// - **on**: enabled
	//
	// - **off**: disabled
	//
	// example:
	//
	// on
	Ssh *string `json:"Ssh,omitempty" xml:"Ssh,omitempty"`
}

func (s DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) String() string {
	return dara.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) GetRdp() *string {
	return s.Rdp
}

func (s *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) GetSqlServer() *string {
	return s.SqlServer
}

func (s *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) GetSsh() *string {
	return s.Ssh
}

func (s *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) SetRdp(v string) *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType {
	s.Rdp = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) SetSqlServer(v string) *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType {
	s.SqlServer = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) SetSsh(v string) *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType {
	s.Ssh = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRulesProtocolType) Validate() error {
	return dara.Validate(s)
}
