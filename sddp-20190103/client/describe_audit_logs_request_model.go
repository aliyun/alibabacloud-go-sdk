// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncRequestId(v string) *DescribeAuditLogsRequest
	GetAsyncRequestId() *string
	SetClientIp(v string) *DescribeAuditLogsRequest
	GetClientIp() *string
	SetClientUa(v string) *DescribeAuditLogsRequest
	GetClientUa() *string
	SetCurrentPage(v int32) *DescribeAuditLogsRequest
	GetCurrentPage() *int32
	SetDatabaseName(v string) *DescribeAuditLogsRequest
	GetDatabaseName() *string
	SetEffectRowRange(v string) *DescribeAuditLogsRequest
	GetEffectRowRange() *string
	SetEndTime(v int64) *DescribeAuditLogsRequest
	GetEndTime() *int64
	SetExecuteTimeRange(v string) *DescribeAuditLogsRequest
	GetExecuteTimeRange() *string
	SetInstanceName(v string) *DescribeAuditLogsRequest
	GetInstanceName() *string
	SetIpType(v string) *DescribeAuditLogsRequest
	GetIpType() *string
	SetLang(v string) *DescribeAuditLogsRequest
	GetLang() *string
	SetLoadWhiteList(v bool) *DescribeAuditLogsRequest
	GetLoadWhiteList() *bool
	SetLogSource(v string) *DescribeAuditLogsRequest
	GetLogSource() *string
	SetMemberAccount(v string) *DescribeAuditLogsRequest
	GetMemberAccount() *string
	SetMessage(v string) *DescribeAuditLogsRequest
	GetMessage() *string
	SetOperateType(v string) *DescribeAuditLogsRequest
	GetOperateType() *string
	SetOssObjectKey(v string) *DescribeAuditLogsRequest
	GetOssObjectKey() *string
	SetPageSize(v int32) *DescribeAuditLogsRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeAuditLogsRequest
	GetProductCode() *string
	SetProductId(v int64) *DescribeAuditLogsRequest
	GetProductId() *int64
	SetRuleAggQuery(v bool) *DescribeAuditLogsRequest
	GetRuleAggQuery() *bool
	SetRuleCategory(v string) *DescribeAuditLogsRequest
	GetRuleCategory() *string
	SetRuleID(v string) *DescribeAuditLogsRequest
	GetRuleID() *string
	SetRuleName(v string) *DescribeAuditLogsRequest
	GetRuleName() *string
	SetSqlText(v string) *DescribeAuditLogsRequest
	GetSqlText() *string
	SetStartTime(v int64) *DescribeAuditLogsRequest
	GetStartTime() *int64
	SetUserName(v string) *DescribeAuditLogsRequest
	GetUserName() *string
}

type DescribeAuditLogsRequest struct {
	// example:
	//
	// CE4681BA-8019-5BE1-9F4B-8973BEA9CF57
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" xml:"AsyncRequestId,omitempty"`
	// example:
	//
	// 10.*.*.94
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// mysql
	ClientUa *string `json:"ClientUa,omitempty" xml:"ClientUa,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// bose_o2o_data
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// in[1 33]
	EffectRowRange *string `json:"EffectRowRange,omitempty" xml:"EffectRowRange,omitempty"`
	// example:
	//
	// 1583856000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// in[1000 2000]
	ExecuteTimeRange *string `json:"ExecuteTimeRange,omitempty" xml:"ExecuteTimeRange,omitempty"`
	// example:
	//
	// i-2zeftaaq4gqcqb9kfkzg
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// aliyun
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// false
	LoadWhiteList *bool `json:"LoadWhiteList,omitempty" xml:"LoadWhiteList,omitempty"`
	// example:
	//
	// SLOW_SQL
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// example:
	//
	// **********8103
	MemberAccount *string `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// Insert
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// oss-key
	OssObjectKey *string `json:"OssObjectKey,omitempty" xml:"OssObjectKey,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId    *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RuleAggQuery *bool  `json:"RuleAggQuery,omitempty" xml:"RuleAggQuery,omitempty"`
	// example:
	//
	// 10
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	// example:
	//
	// 994007
	RuleID   *string `json:"RuleID,omitempty" xml:"RuleID,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// select 	- from test03
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// example:
	//
	// 1616068534877
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// admin
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeAuditLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogsRequest) GetAsyncRequestId() *string {
	return s.AsyncRequestId
}

func (s *DescribeAuditLogsRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeAuditLogsRequest) GetClientUa() *string {
	return s.ClientUa
}

func (s *DescribeAuditLogsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAuditLogsRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeAuditLogsRequest) GetEffectRowRange() *string {
	return s.EffectRowRange
}

func (s *DescribeAuditLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAuditLogsRequest) GetExecuteTimeRange() *string {
	return s.ExecuteTimeRange
}

func (s *DescribeAuditLogsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAuditLogsRequest) GetIpType() *string {
	return s.IpType
}

func (s *DescribeAuditLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuditLogsRequest) GetLoadWhiteList() *bool {
	return s.LoadWhiteList
}

func (s *DescribeAuditLogsRequest) GetLogSource() *string {
	return s.LogSource
}

func (s *DescribeAuditLogsRequest) GetMemberAccount() *string {
	return s.MemberAccount
}

func (s *DescribeAuditLogsRequest) GetMessage() *string {
	return s.Message
}

func (s *DescribeAuditLogsRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *DescribeAuditLogsRequest) GetOssObjectKey() *string {
	return s.OssObjectKey
}

func (s *DescribeAuditLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAuditLogsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeAuditLogsRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeAuditLogsRequest) GetRuleAggQuery() *bool {
	return s.RuleAggQuery
}

func (s *DescribeAuditLogsRequest) GetRuleCategory() *string {
	return s.RuleCategory
}

func (s *DescribeAuditLogsRequest) GetRuleID() *string {
	return s.RuleID
}

func (s *DescribeAuditLogsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAuditLogsRequest) GetSqlText() *string {
	return s.SqlText
}

func (s *DescribeAuditLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAuditLogsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeAuditLogsRequest) SetAsyncRequestId(v string) *DescribeAuditLogsRequest {
	s.AsyncRequestId = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetClientIp(v string) *DescribeAuditLogsRequest {
	s.ClientIp = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetClientUa(v string) *DescribeAuditLogsRequest {
	s.ClientUa = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetCurrentPage(v int32) *DescribeAuditLogsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetDatabaseName(v string) *DescribeAuditLogsRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetEffectRowRange(v string) *DescribeAuditLogsRequest {
	s.EffectRowRange = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetEndTime(v int64) *DescribeAuditLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetExecuteTimeRange(v string) *DescribeAuditLogsRequest {
	s.ExecuteTimeRange = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetInstanceName(v string) *DescribeAuditLogsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetIpType(v string) *DescribeAuditLogsRequest {
	s.IpType = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetLang(v string) *DescribeAuditLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetLoadWhiteList(v bool) *DescribeAuditLogsRequest {
	s.LoadWhiteList = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetLogSource(v string) *DescribeAuditLogsRequest {
	s.LogSource = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetMemberAccount(v string) *DescribeAuditLogsRequest {
	s.MemberAccount = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetMessage(v string) *DescribeAuditLogsRequest {
	s.Message = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetOperateType(v string) *DescribeAuditLogsRequest {
	s.OperateType = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetOssObjectKey(v string) *DescribeAuditLogsRequest {
	s.OssObjectKey = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetPageSize(v int32) *DescribeAuditLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetProductCode(v string) *DescribeAuditLogsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetProductId(v int64) *DescribeAuditLogsRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetRuleAggQuery(v bool) *DescribeAuditLogsRequest {
	s.RuleAggQuery = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetRuleCategory(v string) *DescribeAuditLogsRequest {
	s.RuleCategory = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetRuleID(v string) *DescribeAuditLogsRequest {
	s.RuleID = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetRuleName(v string) *DescribeAuditLogsRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetSqlText(v string) *DescribeAuditLogsRequest {
	s.SqlText = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetStartTime(v int64) *DescribeAuditLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditLogsRequest) SetUserName(v string) *DescribeAuditLogsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeAuditLogsRequest) Validate() error {
	return dara.Validate(s)
}
