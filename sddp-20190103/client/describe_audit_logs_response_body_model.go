// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncRequestId(v string) *DescribeAuditLogsResponseBody
	GetAsyncRequestId() *string
	SetCurrentPage(v int32) *DescribeAuditLogsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeAuditLogsResponseBodyItems) *DescribeAuditLogsResponseBody
	GetItems() []*DescribeAuditLogsResponseBodyItems
	SetPageSize(v int32) *DescribeAuditLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAuditLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAuditLogsResponseBody
	GetTotalCount() *int32
}

type DescribeAuditLogsResponseBody struct {
	// example:
	//
	// CE4681BA-8019-5BE1-9F4B-8973BEA9CF57
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" xml:"AsyncRequestId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribeAuditLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A7021857-AFD9-5AD6-979D-AA7DFC5AFADF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogsResponseBody) GetAsyncRequestId() *string {
	return s.AsyncRequestId
}

func (s *DescribeAuditLogsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAuditLogsResponseBody) GetItems() []*DescribeAuditLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeAuditLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAuditLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAuditLogsResponseBody) SetAsyncRequestId(v string) *DescribeAuditLogsResponseBody {
	s.AsyncRequestId = &v
	return s
}

func (s *DescribeAuditLogsResponseBody) SetCurrentPage(v int32) *DescribeAuditLogsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditLogsResponseBody) SetItems(v []*DescribeAuditLogsResponseBodyItems) *DescribeAuditLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAuditLogsResponseBody) SetPageSize(v int32) *DescribeAuditLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogsResponseBody) SetRequestId(v string) *DescribeAuditLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogsResponseBody) SetTotalCount(v int32) *DescribeAuditLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAuditLogsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAuditLogsResponseBodyItems struct {
	// example:
	//
	// 139.*.*.57
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 35756
	ClientPort *string `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	// example:
	//
	// mysql
	ClientUa *string `json:"ClientUa,omitempty" xml:"ClientUa,omitempty"`
	// example:
	//
	// hide14
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// 1536751124000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// plan_id ~ application_id ~ plan_type ~ plan_name ~ plan_value_id
	DataSet *string `json:"DataSet,omitempty" xml:"DataSet,omitempty"`
	// example:
	//
	// chngc_b2b_migration_sh
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DbType       *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// 10
	EffectRow *int64 `json:"EffectRow,omitempty" xml:"EffectRow,omitempty"`
	// example:
	//
	// 1
	ExecuteStatus *int32 `json:"ExecuteStatus,omitempty" xml:"ExecuteStatus,omitempty"`
	// example:
	//
	// 1751870592000
	ExecuteTime *int64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// example:
	//
	// true
	InWhiteList *bool `json:"InWhiteList,omitempty" xml:"InWhiteList,omitempty"`
	// example:
	//
	// 2
	InstanceAuditStatus *string `json:"InstanceAuditStatus,omitempty" xml:"InstanceAuditStatus,omitempty"`
	// example:
	//
	// instance dscription
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// example:
	//
	// rm-1234
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// aliyun
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// example:
	//
	// SLOW_SQL
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// example:
	//
	// 1751870592000
	LogTime *int64 `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// example:
	//
	// **********8103
	MemberAccount *string `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Drop
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// oss-key1
	OssObjectKey *string `json:"OssObjectKey,omitempty" xml:"OssObjectKey,omitempty"`
	// example:
	//
	// com.sinosoft.chinalife
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
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
	// 10
	RuleCategory *string `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	// example:
	//
	// 9952275
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// select schema_name, catalog_name, default_character_set_name as encoding from information_schema.schemata  order by 1
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// example:
	//
	// it_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 19********94
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// admin
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// 2
	WarnLevel     *string `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
	WarnLevelName *string `json:"WarnLevelName,omitempty" xml:"WarnLevelName,omitempty"`
}

func (s DescribeAuditLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogsResponseBodyItems) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeAuditLogsResponseBodyItems) GetClientPort() *string {
	return s.ClientPort
}

func (s *DescribeAuditLogsResponseBodyItems) GetClientUa() *string {
	return s.ClientUa
}

func (s *DescribeAuditLogsResponseBodyItems) GetColumnName() *string {
	return s.ColumnName
}

func (s *DescribeAuditLogsResponseBodyItems) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeAuditLogsResponseBodyItems) GetDataSet() *string {
	return s.DataSet
}

func (s *DescribeAuditLogsResponseBodyItems) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeAuditLogsResponseBodyItems) GetDbType() *string {
	return s.DbType
}

func (s *DescribeAuditLogsResponseBodyItems) GetEffectRow() *int64 {
	return s.EffectRow
}

func (s *DescribeAuditLogsResponseBodyItems) GetExecuteStatus() *int32 {
	return s.ExecuteStatus
}

func (s *DescribeAuditLogsResponseBodyItems) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *DescribeAuditLogsResponseBodyItems) GetInWhiteList() *bool {
	return s.InWhiteList
}

func (s *DescribeAuditLogsResponseBodyItems) GetInstanceAuditStatus() *string {
	return s.InstanceAuditStatus
}

func (s *DescribeAuditLogsResponseBodyItems) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeAuditLogsResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAuditLogsResponseBodyItems) GetIpType() *string {
	return s.IpType
}

func (s *DescribeAuditLogsResponseBodyItems) GetLogSource() *string {
	return s.LogSource
}

func (s *DescribeAuditLogsResponseBodyItems) GetLogTime() *int64 {
	return s.LogTime
}

func (s *DescribeAuditLogsResponseBodyItems) GetMemberAccount() *string {
	return s.MemberAccount
}

func (s *DescribeAuditLogsResponseBodyItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeAuditLogsResponseBodyItems) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeAuditLogsResponseBodyItems) GetOperateType() *string {
	return s.OperateType
}

func (s *DescribeAuditLogsResponseBodyItems) GetOssObjectKey() *string {
	return s.OssObjectKey
}

func (s *DescribeAuditLogsResponseBodyItems) GetPackageName() *string {
	return s.PackageName
}

func (s *DescribeAuditLogsResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeAuditLogsResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeAuditLogsResponseBodyItems) GetRuleCategory() *string {
	return s.RuleCategory
}

func (s *DescribeAuditLogsResponseBodyItems) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAuditLogsResponseBodyItems) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAuditLogsResponseBodyItems) GetSqlText() *string {
	return s.SqlText
}

func (s *DescribeAuditLogsResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAuditLogsResponseBodyItems) GetUserId() *string {
	return s.UserId
}

func (s *DescribeAuditLogsResponseBodyItems) GetUserName() *string {
	return s.UserName
}

func (s *DescribeAuditLogsResponseBodyItems) GetWarnLevel() *string {
	return s.WarnLevel
}

func (s *DescribeAuditLogsResponseBodyItems) GetWarnLevelName() *string {
	return s.WarnLevelName
}

func (s *DescribeAuditLogsResponseBodyItems) SetClientIp(v string) *DescribeAuditLogsResponseBodyItems {
	s.ClientIp = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetClientPort(v string) *DescribeAuditLogsResponseBodyItems {
	s.ClientPort = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetClientUa(v string) *DescribeAuditLogsResponseBodyItems {
	s.ClientUa = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetColumnName(v string) *DescribeAuditLogsResponseBodyItems {
	s.ColumnName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetCreationTime(v int64) *DescribeAuditLogsResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetDataSet(v string) *DescribeAuditLogsResponseBodyItems {
	s.DataSet = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetDatabaseName(v string) *DescribeAuditLogsResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetDbType(v string) *DescribeAuditLogsResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetEffectRow(v int64) *DescribeAuditLogsResponseBodyItems {
	s.EffectRow = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetExecuteStatus(v int32) *DescribeAuditLogsResponseBodyItems {
	s.ExecuteStatus = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetExecuteTime(v int64) *DescribeAuditLogsResponseBodyItems {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetInWhiteList(v bool) *DescribeAuditLogsResponseBodyItems {
	s.InWhiteList = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetInstanceAuditStatus(v string) *DescribeAuditLogsResponseBodyItems {
	s.InstanceAuditStatus = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetInstanceDescription(v string) *DescribeAuditLogsResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetInstanceName(v string) *DescribeAuditLogsResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetIpType(v string) *DescribeAuditLogsResponseBodyItems {
	s.IpType = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetLogSource(v string) *DescribeAuditLogsResponseBodyItems {
	s.LogSource = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetLogTime(v int64) *DescribeAuditLogsResponseBodyItems {
	s.LogTime = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetMemberAccount(v string) *DescribeAuditLogsResponseBodyItems {
	s.MemberAccount = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetMessage(v string) *DescribeAuditLogsResponseBodyItems {
	s.Message = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetModelName(v string) *DescribeAuditLogsResponseBodyItems {
	s.ModelName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetOperateType(v string) *DescribeAuditLogsResponseBodyItems {
	s.OperateType = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetOssObjectKey(v string) *DescribeAuditLogsResponseBodyItems {
	s.OssObjectKey = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetPackageName(v string) *DescribeAuditLogsResponseBodyItems {
	s.PackageName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetProductCode(v string) *DescribeAuditLogsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetProductId(v int64) *DescribeAuditLogsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetRuleCategory(v string) *DescribeAuditLogsResponseBodyItems {
	s.RuleCategory = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetRuleId(v string) *DescribeAuditLogsResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetRuleName(v string) *DescribeAuditLogsResponseBodyItems {
	s.RuleName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetSqlText(v string) *DescribeAuditLogsResponseBodyItems {
	s.SqlText = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetTableName(v string) *DescribeAuditLogsResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetUserId(v string) *DescribeAuditLogsResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetUserName(v string) *DescribeAuditLogsResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetWarnLevel(v string) *DescribeAuditLogsResponseBodyItems {
	s.WarnLevel = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) SetWarnLevelName(v string) *DescribeAuditLogsResponseBodyItems {
	s.WarnLevelName = &v
	return s
}

func (s *DescribeAuditLogsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
