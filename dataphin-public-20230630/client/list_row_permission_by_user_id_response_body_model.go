// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRowPermissionByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRowPermissionByUserIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListRowPermissionByUserIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRowPermissionByUserIdResponseBody
	GetMessage() *string
	SetPageResult(v *ListRowPermissionByUserIdResponseBodyPageResult) *ListRowPermissionByUserIdResponseBody
	GetPageResult() *ListRowPermissionByUserIdResponseBodyPageResult
	SetRequestId(v string) *ListRowPermissionByUserIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRowPermissionByUserIdResponseBody
	GetSuccess() *bool
}

type ListRowPermissionByUserIdResponseBody struct {
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
	Message    *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListRowPermissionByUserIdResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRowPermissionByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRowPermissionByUserIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRowPermissionByUserIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRowPermissionByUserIdResponseBody) GetPageResult() *ListRowPermissionByUserIdResponseBodyPageResult {
	return s.PageResult
}

func (s *ListRowPermissionByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRowPermissionByUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRowPermissionByUserIdResponseBody) SetCode(v string) *ListRowPermissionByUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBody) SetHttpStatusCode(v int32) *ListRowPermissionByUserIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBody) SetMessage(v string) *ListRowPermissionByUserIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBody) SetPageResult(v *ListRowPermissionByUserIdResponseBodyPageResult) *ListRowPermissionByUserIdResponseBody {
	s.PageResult = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBody) SetRequestId(v string) *ListRowPermissionByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBody) SetSuccess(v bool) *ListRowPermissionByUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionByUserIdResponseBodyPageResult struct {
	Data []*ListRowPermissionByUserIdResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 121
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRowPermissionByUserIdResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponseBodyPageResult) GetData() []*ListRowPermissionByUserIdResponseBodyPageResultData {
	return s.Data
}

func (s *ListRowPermissionByUserIdResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRowPermissionByUserIdResponseBodyPageResult) SetData(v []*ListRowPermissionByUserIdResponseBodyPageResultData) *ListRowPermissionByUserIdResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResult) SetTotalCount(v int32) *ListRowPermissionByUserIdResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResult) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionByUserIdResponseBodyPageResultData struct {
	// example:
	//
	// 30008888
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2023-03-30T21:37:23Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-03-03T10:14Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 30008888
	Modifier *string                                                      `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	Rules    []*ListRowPermissionByUserIdResponseBodyPageResultDataRules  `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Tables   []*ListRowPermissionByUserIdResponseBodyPageResultDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 30000001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListRowPermissionByUserIdResponseBodyPageResultData) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) GetCreator() *string {
	return s.Creator
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) GetModifier() *string {
	return s.Modifier
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) GetRules() []*ListRowPermissionByUserIdResponseBodyPageResultDataRules {
	return s.Rules
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) GetTables() []*ListRowPermissionByUserIdResponseBodyPageResultDataTables {
	return s.Tables
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) SetCreator(v string) *ListRowPermissionByUserIdResponseBodyPageResultData {
	s.Creator = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) SetGmtCreate(v string) *ListRowPermissionByUserIdResponseBodyPageResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) SetGmtModified(v string) *ListRowPermissionByUserIdResponseBodyPageResultData {
	s.GmtModified = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) SetModifier(v string) *ListRowPermissionByUserIdResponseBodyPageResultData {
	s.Modifier = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) SetRules(v []*ListRowPermissionByUserIdResponseBodyPageResultDataRules) *ListRowPermissionByUserIdResponseBodyPageResultData {
	s.Rules = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) SetTables(v []*ListRowPermissionByUserIdResponseBodyPageResultDataTables) *ListRowPermissionByUserIdResponseBodyPageResultData {
	s.Tables = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) SetTenantId(v int64) *ListRowPermissionByUserIdResponseBodyPageResultData {
	s.TenantId = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultData) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionByUserIdResponseBodyPageResultDataRules struct {
	Expressions []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	IsDelete *bool   `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// SELECT_COLUMN
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// example:
	//
	// 1
	Status          *int32                                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	UserMappingList []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList `json:"UserMappingList,omitempty" xml:"UserMappingList,omitempty" type:"Repeated"`
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataRules) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataRules) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) GetExpressions() []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions {
	return s.Expressions
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) GetIsDelete() *bool {
	return s.IsDelete
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) GetStatus() *int32 {
	return s.Status
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) GetUserMappingList() []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList {
	return s.UserMappingList
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) SetExpressions(v []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) *ListRowPermissionByUserIdResponseBodyPageResultDataRules {
	s.Expressions = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) SetIsDelete(v bool) *ListRowPermissionByUserIdResponseBodyPageResultDataRules {
	s.IsDelete = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) SetRuleName(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataRules {
	s.RuleName = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) SetScopeType(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataRules {
	s.ScopeType = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) SetStatus(v int32) *ListRowPermissionByUserIdResponseBodyPageResultDataRules {
	s.Status = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) SetUserMappingList(v []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList) *ListRowPermissionByUserIdResponseBodyPageResultDataRules {
	s.UserMappingList = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions struct {
	// example:
	//
	// business_id
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// example:
	//
	// EQUAL
	Operator      *string       `json:"Operator,omitempty" xml:"Operator,omitempty"`
	SubConditions []interface{} `json:"SubConditions,omitempty" xml:"SubConditions,omitempty" type:"Repeated"`
	// example:
	//
	// EXPRESSION
	Type   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) GetMappingColumnName() *string {
	return s.MappingColumnName
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) GetOperator() *string {
	return s.Operator
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) GetSubConditions() []interface{} {
	return s.SubConditions
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) GetType() *string {
	return s.Type
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) GetValues() []*string {
	return s.Values
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) SetMappingColumnName(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions {
	s.MappingColumnName = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) SetOperator(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions {
	s.Operator = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) SetSubConditions(v []interface{}) *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions {
	s.SubConditions = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) SetType(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions {
	s.Type = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) SetValues(v []*string) *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions {
	s.Values = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList struct {
	// example:
	//
	// PERSONAL
	AccountType *string                                                                            `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Accounts    []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList) GetAccountType() *string {
	return s.AccountType
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList) GetAccounts() []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts {
	return s.Accounts
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList) SetAccountType(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList {
	s.AccountType = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList) SetAccounts(v []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts) *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList {
	s.Accounts = v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingList) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts struct {
	// example:
	//
	// 30008888
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts) SetAccountId(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts {
	s.AccountId = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts) Validate() error {
	return dara.Validate(s)
}

type ListRowPermissionByUserIdResponseBodyPageResultDataTables struct {
	// example:
	//
	// business_id
	ColumnName        *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// example:
	//
	// odps.300199897.project_name.table_name
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataTables) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdResponseBodyPageResultDataTables) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataTables) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataTables) GetMappingColumnName() *string {
	return s.MappingColumnName
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataTables) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataTables) SetColumnName(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataTables {
	s.ColumnName = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataTables) SetMappingColumnName(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataTables {
	s.MappingColumnName = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataTables) SetResourceId(v string) *ListRowPermissionByUserIdResponseBodyPageResultDataTables {
	s.ResourceId = &v
	return s
}

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataTables) Validate() error {
	return dara.Validate(s)
}
