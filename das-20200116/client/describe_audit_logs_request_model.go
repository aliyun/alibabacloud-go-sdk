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
	SetRuleId(v string) *DescribeAuditLogsRequest
	GetRuleId() *string
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
	// 2c548e83-1473-4fda-b3dc-5a189074ead5
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" xml:"AsyncRequestId,omitempty"`
	// example:
	//
	// 11.26.118.7
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientUa *string `json:"ClientUa,omitempty" xml:"ClientUa,omitempty"`
	// example:
	//
	// 3
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// TestDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// in[1 33]
	EffectRowRange *string `json:"EffectRowRange,omitempty" xml:"EffectRowRange,omitempty"`
	// example:
	//
	// 15682887991222
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// in[1000 2000]
	ExecuteTimeRange *string `json:"ExecuteTimeRange,omitempty" xml:"ExecuteTimeRange,omitempty"`
	// example:
	//
	// rm-t4ni1cezz5y3xxxx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// aliyun
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// example:
	//
	// zh-CN
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
	// Insert
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// N.A
	OssObjectKey *string `json:"OssObjectKey,omitempty" xml:"OssObjectKey,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// true
	RuleAggQuery *bool `json:"RuleAggQuery,omitempty" xml:"RuleAggQuery,omitempty"`
	// example:
	//
	// 10
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	// example:
	//
	// 9953411
	RuleID *string `json:"RuleID,omitempty" xml:"RuleID,omitempty"`
	// example:
	//
	// 867028
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// test_rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// SELECT 	- FROM test where name = \\"das\\"
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// example:
	//
	// 1608888296000
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

func (s *DescribeAuditLogsRequest) GetRuleId() *string {
	return s.RuleId
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

func (s *DescribeAuditLogsRequest) SetRuleId(v string) *DescribeAuditLogsRequest {
	s.RuleId = &v
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
