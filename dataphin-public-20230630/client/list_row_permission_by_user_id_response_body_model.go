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
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result.
	PageResult *ListRowPermissionByUserIdResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRowPermissionByUserIdResponseBodyPageResult struct {
	// The query result.
	Data []*ListRowPermissionByUserIdResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of records.
	//
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
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRowPermissionByUserIdResponseBodyPageResultData struct {
	// The creator.
	//
	// example:
	//
	// 30008888
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2023-03-30T21:37:23Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-03-03T10:14Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The modifier.
	//
	// example:
	//
	// 30008888
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The rules.
	Rules []*ListRowPermissionByUserIdResponseBodyPageResultDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The related tables.
	Tables []*ListRowPermissionByUserIdResponseBodyPageResultDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The tenant ID.
	//
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
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRowPermissionByUserIdResponseBodyPageResultDataRules struct {
	// The rule expressions.
	Expressions []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the rule is deleted.
	//
	// example:
	//
	// 0
	IsDelete *bool `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// The rule name.
	//
	// example:
	//
	// 业务管控
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The scope type of the rule.
	//
	// example:
	//
	// SELECT_COLUMN
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The rule status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The accounts bound to the rule.
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

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) GetId() *int64 {
	return s.Id
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

func (s *ListRowPermissionByUserIdResponseBodyPageResultDataRules) SetId(v int64) *ListRowPermissionByUserIdResponseBodyPageResultDataRules {
	s.Id = &v
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
	if s.Expressions != nil {
		for _, item := range s.Expressions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserMappingList != nil {
		for _, item := range s.UserMappingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRowPermissionByUserIdResponseBodyPageResultDataRulesExpressions struct {
	// The mapping field name.
	//
	// example:
	//
	// business_id
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// The expression operator.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The sub-expressions.
	SubConditions []interface{} `json:"SubConditions,omitempty" xml:"SubConditions,omitempty" type:"Repeated"`
	// The expression type.
	//
	// example:
	//
	// EXPRESSION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The expression operation values.
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
	// The type of the account bound to the rule.
	//
	// example:
	//
	// PERSONAL
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The accounts bound to the rule.
	Accounts []*ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
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
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRowPermissionByUserIdResponseBodyPageResultDataRulesUserMappingListAccounts struct {
	// The ID of the account bound to the rule.
	//
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
	// The table field.
	//
	// example:
	//
	// business_id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The mapping field name.
	//
	// example:
	//
	// 拦截规则
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// The table GUID.
	//
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
