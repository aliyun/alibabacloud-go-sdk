// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRowPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRowPermissionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListRowPermissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRowPermissionResponseBody
	GetMessage() *string
	SetPageResult(v *ListRowPermissionResponseBodyPageResult) *ListRowPermissionResponseBody
	GetPageResult() *ListRowPermissionResponseBodyPageResult
	SetRequestId(v string) *ListRowPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRowPermissionResponseBody
	GetSuccess() *bool
}

type ListRowPermissionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListRowPermissionResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRowPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRowPermissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRowPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRowPermissionResponseBody) GetPageResult() *ListRowPermissionResponseBodyPageResult {
	return s.PageResult
}

func (s *ListRowPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRowPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRowPermissionResponseBody) SetCode(v string) *ListRowPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *ListRowPermissionResponseBody) SetHttpStatusCode(v int32) *ListRowPermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRowPermissionResponseBody) SetMessage(v string) *ListRowPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *ListRowPermissionResponseBody) SetPageResult(v *ListRowPermissionResponseBodyPageResult) *ListRowPermissionResponseBody {
	s.PageResult = v
	return s
}

func (s *ListRowPermissionResponseBody) SetRequestId(v string) *ListRowPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRowPermissionResponseBody) SetSuccess(v bool) *ListRowPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *ListRowPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionResponseBodyPageResult struct {
	Data []*ListRowPermissionResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRowPermissionResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBodyPageResult) GetData() []*ListRowPermissionResponseBodyPageResultData {
	return s.Data
}

func (s *ListRowPermissionResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRowPermissionResponseBodyPageResult) SetData(v []*ListRowPermissionResponseBodyPageResultData) *ListRowPermissionResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResult) SetTotalCount(v int32) *ListRowPermissionResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResult) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionResponseBodyPageResultData struct {
	// example:
	//
	// 30008888
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2023-08-23T08:01:44Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-02-12T02:16:45Z
	GmtModified    *string                                                      `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MappingColumns []*ListRowPermissionResponseBodyPageResultDataMappingColumns `json:"MappingColumns,omitempty" xml:"MappingColumns,omitempty" type:"Repeated"`
	// example:
	//
	// 30008888
	Modifier          *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	RowPermissionDesc *string `json:"RowPermissionDesc,omitempty" xml:"RowPermissionDesc,omitempty"`
	// example:
	//
	// 30008888
	RowPermissionId   *int64                                               `json:"RowPermissionId,omitempty" xml:"RowPermissionId,omitempty"`
	RowPermissionName *string                                              `json:"RowPermissionName,omitempty" xml:"RowPermissionName,omitempty"`
	Rules             []*ListRowPermissionResponseBodyPageResultDataRules  `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Tables            []*ListRowPermissionResponseBodyPageResultDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 30000001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListRowPermissionResponseBodyPageResultData) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBodyPageResultData) GetCreator() *string {
	return s.Creator
}

func (s *ListRowPermissionResponseBodyPageResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListRowPermissionResponseBodyPageResultData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListRowPermissionResponseBodyPageResultData) GetMappingColumns() []*ListRowPermissionResponseBodyPageResultDataMappingColumns {
	return s.MappingColumns
}

func (s *ListRowPermissionResponseBodyPageResultData) GetModifier() *string {
	return s.Modifier
}

func (s *ListRowPermissionResponseBodyPageResultData) GetRowPermissionDesc() *string {
	return s.RowPermissionDesc
}

func (s *ListRowPermissionResponseBodyPageResultData) GetRowPermissionId() *int64 {
	return s.RowPermissionId
}

func (s *ListRowPermissionResponseBodyPageResultData) GetRowPermissionName() *string {
	return s.RowPermissionName
}

func (s *ListRowPermissionResponseBodyPageResultData) GetRules() []*ListRowPermissionResponseBodyPageResultDataRules {
	return s.Rules
}

func (s *ListRowPermissionResponseBodyPageResultData) GetTables() []*ListRowPermissionResponseBodyPageResultDataTables {
	return s.Tables
}

func (s *ListRowPermissionResponseBodyPageResultData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListRowPermissionResponseBodyPageResultData) SetCreator(v string) *ListRowPermissionResponseBodyPageResultData {
	s.Creator = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetGmtCreate(v string) *ListRowPermissionResponseBodyPageResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetGmtModified(v string) *ListRowPermissionResponseBodyPageResultData {
	s.GmtModified = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetMappingColumns(v []*ListRowPermissionResponseBodyPageResultDataMappingColumns) *ListRowPermissionResponseBodyPageResultData {
	s.MappingColumns = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetModifier(v string) *ListRowPermissionResponseBodyPageResultData {
	s.Modifier = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetRowPermissionDesc(v string) *ListRowPermissionResponseBodyPageResultData {
	s.RowPermissionDesc = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetRowPermissionId(v int64) *ListRowPermissionResponseBodyPageResultData {
	s.RowPermissionId = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetRowPermissionName(v string) *ListRowPermissionResponseBodyPageResultData {
	s.RowPermissionName = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetRules(v []*ListRowPermissionResponseBodyPageResultDataRules) *ListRowPermissionResponseBodyPageResultData {
	s.Rules = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetTables(v []*ListRowPermissionResponseBodyPageResultDataTables) *ListRowPermissionResponseBodyPageResultData {
	s.Tables = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) SetTenantId(v int64) *ListRowPermissionResponseBodyPageResultData {
	s.TenantId = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultData) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionResponseBodyPageResultDataMappingColumns struct {
	ColumnDesc *string `json:"ColumnDesc,omitempty" xml:"ColumnDesc,omitempty"`
	// example:
	//
	// business_id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// string
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
}

func (s ListRowPermissionResponseBodyPageResultDataMappingColumns) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBodyPageResultDataMappingColumns) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBodyPageResultDataMappingColumns) GetColumnDesc() *string {
	return s.ColumnDesc
}

func (s *ListRowPermissionResponseBodyPageResultDataMappingColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListRowPermissionResponseBodyPageResultDataMappingColumns) GetColumnType() *string {
	return s.ColumnType
}

func (s *ListRowPermissionResponseBodyPageResultDataMappingColumns) SetColumnDesc(v string) *ListRowPermissionResponseBodyPageResultDataMappingColumns {
	s.ColumnDesc = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataMappingColumns) SetColumnName(v string) *ListRowPermissionResponseBodyPageResultDataMappingColumns {
	s.ColumnName = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataMappingColumns) SetColumnType(v string) *ListRowPermissionResponseBodyPageResultDataMappingColumns {
	s.ColumnType = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataMappingColumns) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionResponseBodyPageResultDataRules struct {
	Expressions []*ListRowPermissionResponseBodyPageResultDataRulesExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	IsDelete    *bool                                                          `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	RuleName    *string                                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// SELECT_COLUMN
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// example:
	//
	// 1
	Status          *int32                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	UserMappingList []*ListRowPermissionResponseBodyPageResultDataRulesUserMappingList `json:"UserMappingList,omitempty" xml:"UserMappingList,omitempty" type:"Repeated"`
}

func (s ListRowPermissionResponseBodyPageResultDataRules) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBodyPageResultDataRules) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) GetExpressions() []*ListRowPermissionResponseBodyPageResultDataRulesExpressions {
	return s.Expressions
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) GetIsDelete() *bool {
	return s.IsDelete
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) GetStatus() *int32 {
	return s.Status
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) GetUserMappingList() []*ListRowPermissionResponseBodyPageResultDataRulesUserMappingList {
	return s.UserMappingList
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) SetExpressions(v []*ListRowPermissionResponseBodyPageResultDataRulesExpressions) *ListRowPermissionResponseBodyPageResultDataRules {
	s.Expressions = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) SetIsDelete(v bool) *ListRowPermissionResponseBodyPageResultDataRules {
	s.IsDelete = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) SetRuleName(v string) *ListRowPermissionResponseBodyPageResultDataRules {
	s.RuleName = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) SetScopeType(v string) *ListRowPermissionResponseBodyPageResultDataRules {
	s.ScopeType = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) SetStatus(v int32) *ListRowPermissionResponseBodyPageResultDataRules {
	s.Status = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) SetUserMappingList(v []*ListRowPermissionResponseBodyPageResultDataRulesUserMappingList) *ListRowPermissionResponseBodyPageResultDataRules {
	s.UserMappingList = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRules) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionResponseBodyPageResultDataRulesExpressions struct {
	// example:
	//
	// business_id
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// example:
	//
	// 30008888
	Operator      *string       `json:"Operator,omitempty" xml:"Operator,omitempty"`
	SubConditions []interface{} `json:"SubConditions,omitempty" xml:"SubConditions,omitempty" type:"Repeated"`
	// example:
	//
	// EXPRESSION
	Type   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRowPermissionResponseBodyPageResultDataRulesExpressions) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBodyPageResultDataRulesExpressions) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) GetMappingColumnName() *string {
	return s.MappingColumnName
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) GetOperator() *string {
	return s.Operator
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) GetSubConditions() []interface{} {
	return s.SubConditions
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) GetType() *string {
	return s.Type
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) GetValues() []*string {
	return s.Values
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) SetMappingColumnName(v string) *ListRowPermissionResponseBodyPageResultDataRulesExpressions {
	s.MappingColumnName = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) SetOperator(v string) *ListRowPermissionResponseBodyPageResultDataRulesExpressions {
	s.Operator = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) SetSubConditions(v []interface{}) *ListRowPermissionResponseBodyPageResultDataRulesExpressions {
	s.SubConditions = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) SetType(v string) *ListRowPermissionResponseBodyPageResultDataRulesExpressions {
	s.Type = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) SetValues(v []*string) *ListRowPermissionResponseBodyPageResultDataRulesExpressions {
	s.Values = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesExpressions) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionResponseBodyPageResultDataRulesUserMappingList struct {
	// example:
	//
	// PERSONAL
	AccountType *string                                                                    `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Accounts    []*ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s ListRowPermissionResponseBodyPageResultDataRulesUserMappingList) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBodyPageResultDataRulesUserMappingList) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesUserMappingList) GetAccountType() *string {
	return s.AccountType
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesUserMappingList) GetAccounts() []*ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts {
	return s.Accounts
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesUserMappingList) SetAccountType(v string) *ListRowPermissionResponseBodyPageResultDataRulesUserMappingList {
	s.AccountType = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesUserMappingList) SetAccounts(v []*ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts) *ListRowPermissionResponseBodyPageResultDataRulesUserMappingList {
	s.Accounts = v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesUserMappingList) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts struct {
	// example:
	//
	// 30008888
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts) SetAccountId(v string) *ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts {
	s.AccountId = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataRulesUserMappingListAccounts) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionResponseBodyPageResultDataTables struct {
	// example:
	//
	// business_id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// business_id
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// example:
	//
	// odps.300199897.project_name.table_name
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ListRowPermissionResponseBodyPageResultDataTables) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponseBodyPageResultDataTables) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponseBodyPageResultDataTables) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListRowPermissionResponseBodyPageResultDataTables) GetMappingColumnName() *string {
	return s.MappingColumnName
}

func (s *ListRowPermissionResponseBodyPageResultDataTables) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListRowPermissionResponseBodyPageResultDataTables) SetColumnName(v string) *ListRowPermissionResponseBodyPageResultDataTables {
	s.ColumnName = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataTables) SetMappingColumnName(v string) *ListRowPermissionResponseBodyPageResultDataTables {
	s.MappingColumnName = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataTables) SetResourceId(v string) *ListRowPermissionResponseBodyPageResultDataTables {
	s.ResourceId = &v
	return s
}

func (s *ListRowPermissionResponseBodyPageResultDataTables) Validate() error {
	return dara.Validate(s)
}
