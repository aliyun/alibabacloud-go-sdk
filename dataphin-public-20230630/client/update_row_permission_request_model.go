// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRowPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateRowPermissionRequest
	GetOpTenantId() *int64
	SetUpdateRowPermissionCommand(v *UpdateRowPermissionRequestUpdateRowPermissionCommand) *UpdateRowPermissionRequest
	GetUpdateRowPermissionCommand() *UpdateRowPermissionRequestUpdateRowPermissionCommand
}

type UpdateRowPermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateRowPermissionCommand *UpdateRowPermissionRequestUpdateRowPermissionCommand `json:"UpdateRowPermissionCommand,omitempty" xml:"UpdateRowPermissionCommand,omitempty" type:"Struct"`
}

func (s UpdateRowPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateRowPermissionRequest) GetUpdateRowPermissionCommand() *UpdateRowPermissionRequestUpdateRowPermissionCommand {
	return s.UpdateRowPermissionCommand
}

func (s *UpdateRowPermissionRequest) SetOpTenantId(v int64) *UpdateRowPermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateRowPermissionRequest) SetUpdateRowPermissionCommand(v *UpdateRowPermissionRequestUpdateRowPermissionCommand) *UpdateRowPermissionRequest {
	s.UpdateRowPermissionCommand = v
	return s
}

func (s *UpdateRowPermissionRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRowPermissionRequestUpdateRowPermissionCommand struct {
	// This parameter is required.
	MappingColumns    []*UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns `json:"MappingColumns,omitempty" xml:"MappingColumns,omitempty" type:"Repeated"`
	RowPermissionDesc *string                                                               `json:"RowPermissionDesc,omitempty" xml:"RowPermissionDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30009999
	RowPermissionId *int64 `json:"RowPermissionId,omitempty" xml:"RowPermissionId,omitempty"`
	// This parameter is required.
	RowPermissionName *string                                                       `json:"RowPermissionName,omitempty" xml:"RowPermissionName,omitempty"`
	Rules             []*UpdateRowPermissionRequestUpdateRowPermissionCommandRules  `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Tables            []*UpdateRowPermissionRequestUpdateRowPermissionCommandTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommand) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) GetMappingColumns() []*UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns {
	return s.MappingColumns
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) GetRowPermissionDesc() *string {
	return s.RowPermissionDesc
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) GetRowPermissionId() *int64 {
	return s.RowPermissionId
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) GetRowPermissionName() *string {
	return s.RowPermissionName
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) GetRules() []*UpdateRowPermissionRequestUpdateRowPermissionCommandRules {
	return s.Rules
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) GetTables() []*UpdateRowPermissionRequestUpdateRowPermissionCommandTables {
	return s.Tables
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) SetMappingColumns(v []*UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) *UpdateRowPermissionRequestUpdateRowPermissionCommand {
	s.MappingColumns = v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) SetRowPermissionDesc(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommand {
	s.RowPermissionDesc = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) SetRowPermissionId(v int64) *UpdateRowPermissionRequestUpdateRowPermissionCommand {
	s.RowPermissionId = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) SetRowPermissionName(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommand {
	s.RowPermissionName = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) SetRules(v []*UpdateRowPermissionRequestUpdateRowPermissionCommandRules) *UpdateRowPermissionRequestUpdateRowPermissionCommand {
	s.Rules = v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) SetTables(v []*UpdateRowPermissionRequestUpdateRowPermissionCommandTables) *UpdateRowPermissionRequestUpdateRowPermissionCommand {
	s.Tables = v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommand) Validate() error {
	return dara.Validate(s)
}

type UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns struct {
	ColumnDesc *string `json:"ColumnDesc,omitempty" xml:"ColumnDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// business_id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRING
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) GetColumnDesc() *string {
	return s.ColumnDesc
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) GetColumnType() *string {
	return s.ColumnType
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) SetColumnDesc(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns {
	s.ColumnDesc = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) SetColumnName(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns {
	s.ColumnName = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) SetColumnType(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns {
	s.ColumnType = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandMappingColumns) Validate() error {
	return dara.Validate(s)
}

type UpdateRowPermissionRequestUpdateRowPermissionCommandRules struct {
	// This parameter is required.
	Expressions []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	IsDelete *bool `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SELECT_COLUMN
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status          *int32                                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	UserMappingList []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList `json:"UserMappingList,omitempty" xml:"UserMappingList,omitempty" type:"Repeated"`
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandRules) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) GetExpressions() []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions {
	return s.Expressions
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) GetIsDelete() *bool {
	return s.IsDelete
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) GetScopeType() *string {
	return s.ScopeType
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) GetUserMappingList() []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList {
	return s.UserMappingList
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) SetExpressions(v []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) *UpdateRowPermissionRequestUpdateRowPermissionCommandRules {
	s.Expressions = v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) SetIsDelete(v bool) *UpdateRowPermissionRequestUpdateRowPermissionCommandRules {
	s.IsDelete = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) SetRuleName(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandRules {
	s.RuleName = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) SetScopeType(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandRules {
	s.ScopeType = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) SetStatus(v int32) *UpdateRowPermissionRequestUpdateRowPermissionCommandRules {
	s.Status = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) SetUserMappingList(v []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList) *UpdateRowPermissionRequestUpdateRowPermissionCommandRules {
	s.UserMappingList = v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRules) Validate() error {
	return dara.Validate(s)
}

type UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions struct {
	// This parameter is required.
	//
	// example:
	//
	// business_id
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OR
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	SubConditions []interface{} `json:"SubConditions,omitempty" xml:"SubConditions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) GetMappingColumnName() *string {
	return s.MappingColumnName
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) GetOperator() *string {
	return s.Operator
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) GetSubConditions() []interface{} {
	return s.SubConditions
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) GetType() *string {
	return s.Type
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) GetValues() []*string {
	return s.Values
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) SetMappingColumnName(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions {
	s.MappingColumnName = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) SetOperator(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions {
	s.Operator = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) SetSubConditions(v []interface{}) *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions {
	s.SubConditions = v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) SetType(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions {
	s.Type = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) SetValues(v []*string) *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions {
	s.Values = v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesExpressions) Validate() error {
	return dara.Validate(s)
}

type UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList struct {
	// This parameter is required.
	//
	// example:
	//
	// PERSONAL
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	Accounts []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList) GetAccountType() *string {
	return s.AccountType
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList) GetAccounts() []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts {
	return s.Accounts
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList) SetAccountType(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList {
	s.AccountType = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList) SetAccounts(v []*UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts) *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList {
	s.Accounts = v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingList) Validate() error {
	return dara.Validate(s)
}

type UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts) SetAccountId(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts {
	s.AccountId = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandRulesUserMappingListAccounts) Validate() error {
	return dara.Validate(s)
}

type UpdateRowPermissionRequestUpdateRowPermissionCommandTables struct {
	// This parameter is required.
	//
	// example:
	//
	// business_id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// business_id
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// odps.300199897.project_name.table_name
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandTables) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionRequestUpdateRowPermissionCommandTables) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandTables) GetColumnName() *string {
	return s.ColumnName
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandTables) GetMappingColumnName() *string {
	return s.MappingColumnName
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandTables) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandTables) SetColumnName(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandTables {
	s.ColumnName = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandTables) SetMappingColumnName(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandTables {
	s.MappingColumnName = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandTables) SetResourceId(v string) *UpdateRowPermissionRequestUpdateRowPermissionCommandTables {
	s.ResourceId = &v
	return s
}

func (s *UpdateRowPermissionRequestUpdateRowPermissionCommandTables) Validate() error {
	return dara.Validate(s)
}
