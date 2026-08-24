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
	SetLogQueryOpJson(v string) *DescribeAuditLogsRequest
	GetLogQueryOpJson() *string
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
	// The request ID.
	//
	// example:
	//
	// 2c548e83-1473-4fda-b3dc-5a189074ead5
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" xml:"AsyncRequestId,omitempty"`
	// The IP address of the request client.
	//
	// example:
	//
	// 11.26.118.7
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The client type.
	//
	// example:
	//
	// mysql
	ClientUa *string `json:"ClientUa,omitempty" xml:"ClientUa,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 3
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// TestDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The range of affected rows.
	//
	// example:
	//
	// in[1 33]
	EffectRowRange *string `json:"EffectRowRange,omitempty" xml:"EffectRowRange,omitempty"`
	// The end time of the alert log. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 15682887991222
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The range of execution time.
	//
	// example:
	//
	// in[1000 2000]
	ExecuteTimeRange *string `json:"ExecuteTimeRange,omitempty" xml:"ExecuteTimeRange,omitempty"`
	// The name of the asset instance.
	//
	// example:
	//
	// rm-t4ni1cezz5y3xxxx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network type. Valid values:
	//
	// - **default*	- (default): non-Alibaba Cloud service.
	//
	// - **aliyun**: Alibaba Cloud service.
	//
	// example:
	//
	// aliyun
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The language type for sending and receiving messages. Default value: zh_cn. Valid values:
	//
	// - zh_cn: Chinese.
	//
	// - en_us: English.
	//
	// example:
	//
	// zh-CN
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to load the whitelist status.
	//
	// example:
	//
	// false
	LoadWhiteList  *bool   `json:"LoadWhiteList,omitempty" xml:"LoadWhiteList,omitempty"`
	LogQueryOpJson *string `json:"LogQueryOpJson,omitempty" xml:"LogQueryOpJson,omitempty"`
	// The data source.
	//
	// example:
	//
	// SLOW_SQL
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The operation type.
	//
	// example:
	//
	// Insert
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The key of the OSS storage object.
	//
	// example:
	//
	// N.A
	OssObjectKey *string `json:"OssObjectKey,omitempty" xml:"OssObjectKey,omitempty"`
	// The maximum number of records per page. Default value: 10.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the product to which the data asset belongs. Valid values: **RDS, PolarDB, PolarDB-X 2.0**, and others.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID that corresponds to the product name to which the data object belongs. Valid values:
	//
	// - 3: ADB-MYSQL
	//
	// - 5: RDS
	//
	// - 7: PolarDB-X
	//
	// - 8: PolarDB
	//
	// - 9: ADB-PG
	//
	// - 11: MongoDB
	//
	// - 25: Redis
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// Specifies whether to perform an aggregate query.
	//
	// example:
	//
	// true
	RuleAggQuery *bool `json:"RuleAggQuery,omitempty" xml:"RuleAggQuery,omitempty"`
	// The rule type.
	//
	// example:
	//
	// 10
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	// The ID of the audit policy.
	//
	// example:
	//
	// 9953411
	RuleID *string `json:"RuleID,omitempty" xml:"RuleID,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 867028
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the audit policy.
	//
	// example:
	//
	// test_rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The content of the SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM test where name = \\"das\\"
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// The start time of the alert log, in milliseconds.
	//
	// example:
	//
	// 1608888296000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username.
	//
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

func (s *DescribeAuditLogsRequest) GetLogQueryOpJson() *string {
	return s.LogQueryOpJson
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

func (s *DescribeAuditLogsRequest) SetLogQueryOpJson(v string) *DescribeAuditLogsRequest {
	s.LogQueryOpJson = &v
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
