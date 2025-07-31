// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRowPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateRowPermissionCommand(v *CreateRowPermissionRequestCreateRowPermissionCommand) *CreateRowPermissionRequest
	GetCreateRowPermissionCommand() *CreateRowPermissionRequestCreateRowPermissionCommand
	SetOpTenantId(v int64) *CreateRowPermissionRequest
	GetOpTenantId() *int64
}

type CreateRowPermissionRequest struct {
	// This parameter is required.
	CreateRowPermissionCommand *CreateRowPermissionRequestCreateRowPermissionCommand `json:"CreateRowPermissionCommand,omitempty" xml:"CreateRowPermissionCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateRowPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionRequest) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionRequest) GetCreateRowPermissionCommand() *CreateRowPermissionRequestCreateRowPermissionCommand {
	return s.CreateRowPermissionCommand
}

func (s *CreateRowPermissionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateRowPermissionRequest) SetCreateRowPermissionCommand(v *CreateRowPermissionRequestCreateRowPermissionCommand) *CreateRowPermissionRequest {
	s.CreateRowPermissionCommand = v
	return s
}

func (s *CreateRowPermissionRequest) SetOpTenantId(v int64) *CreateRowPermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateRowPermissionRequest) Validate() error {
	return dara.Validate(s)
}

type CreateRowPermissionRequestCreateRowPermissionCommand struct {
	// This parameter is required.
	MappingColumns    []*CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns `json:"MappingColumns,omitempty" xml:"MappingColumns,omitempty" type:"Repeated"`
	RowPermissionDesc *string                                                               `json:"RowPermissionDesc,omitempty" xml:"RowPermissionDesc,omitempty"`
	// This parameter is required.
	RowPermissionName *string                                                       `json:"RowPermissionName,omitempty" xml:"RowPermissionName,omitempty"`
	Rules             []*CreateRowPermissionRequestCreateRowPermissionCommandRules  `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Tables            []*CreateRowPermissionRequestCreateRowPermissionCommandTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s CreateRowPermissionRequestCreateRowPermissionCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionRequestCreateRowPermissionCommand) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) GetMappingColumns() []*CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns {
	return s.MappingColumns
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) GetRowPermissionDesc() *string {
	return s.RowPermissionDesc
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) GetRowPermissionName() *string {
	return s.RowPermissionName
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) GetRules() []*CreateRowPermissionRequestCreateRowPermissionCommandRules {
	return s.Rules
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) GetTables() []*CreateRowPermissionRequestCreateRowPermissionCommandTables {
	return s.Tables
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) SetMappingColumns(v []*CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) *CreateRowPermissionRequestCreateRowPermissionCommand {
	s.MappingColumns = v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) SetRowPermissionDesc(v string) *CreateRowPermissionRequestCreateRowPermissionCommand {
	s.RowPermissionDesc = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) SetRowPermissionName(v string) *CreateRowPermissionRequestCreateRowPermissionCommand {
	s.RowPermissionName = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) SetRules(v []*CreateRowPermissionRequestCreateRowPermissionCommandRules) *CreateRowPermissionRequestCreateRowPermissionCommand {
	s.Rules = v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) SetTables(v []*CreateRowPermissionRequestCreateRowPermissionCommandTables) *CreateRowPermissionRequestCreateRowPermissionCommand {
	s.Tables = v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommand) Validate() error {
	return dara.Validate(s)
}

type CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns struct {
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

func (s CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) GetColumnDesc() *string {
	return s.ColumnDesc
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) GetColumnType() *string {
	return s.ColumnType
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) SetColumnDesc(v string) *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns {
	s.ColumnDesc = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) SetColumnName(v string) *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns {
	s.ColumnName = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) SetColumnType(v string) *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns {
	s.ColumnType = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandMappingColumns) Validate() error {
	return dara.Validate(s)
}

type CreateRowPermissionRequestCreateRowPermissionCommandRules struct {
	// This parameter is required.
	Expressions []*CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
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
	UserMappingList []*CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList `json:"UserMappingList,omitempty" xml:"UserMappingList,omitempty" type:"Repeated"`
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandRules) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandRules) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) GetExpressions() []*CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions {
	return s.Expressions
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) GetIsDelete() *bool {
	return s.IsDelete
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) GetScopeType() *string {
	return s.ScopeType
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) GetStatus() *int32 {
	return s.Status
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) GetUserMappingList() []*CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList {
	return s.UserMappingList
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) SetExpressions(v []*CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) *CreateRowPermissionRequestCreateRowPermissionCommandRules {
	s.Expressions = v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) SetIsDelete(v bool) *CreateRowPermissionRequestCreateRowPermissionCommandRules {
	s.IsDelete = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) SetRuleName(v string) *CreateRowPermissionRequestCreateRowPermissionCommandRules {
	s.RuleName = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) SetScopeType(v string) *CreateRowPermissionRequestCreateRowPermissionCommandRules {
	s.ScopeType = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) SetStatus(v int32) *CreateRowPermissionRequestCreateRowPermissionCommandRules {
	s.Status = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) SetUserMappingList(v []*CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList) *CreateRowPermissionRequestCreateRowPermissionCommandRules {
	s.UserMappingList = v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRules) Validate() error {
	return dara.Validate(s)
}

type CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions struct {
	// This parameter is required.
	//
	// example:
	//
	// id
	MappingColumnName *string `json:"MappingColumnName,omitempty" xml:"MappingColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	SubConditions []interface{} `json:"SubConditions,omitempty" xml:"SubConditions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// RELATION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) GetMappingColumnName() *string {
	return s.MappingColumnName
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) GetOperator() *string {
	return s.Operator
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) GetSubConditions() []interface{} {
	return s.SubConditions
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) GetType() *string {
	return s.Type
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) GetValues() []*string {
	return s.Values
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) SetMappingColumnName(v string) *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions {
	s.MappingColumnName = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) SetOperator(v string) *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions {
	s.Operator = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) SetSubConditions(v []interface{}) *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions {
	s.SubConditions = v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) SetType(v string) *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions {
	s.Type = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) SetValues(v []*string) *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions {
	s.Values = v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesExpressions) Validate() error {
	return dara.Validate(s)
}

type CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList struct {
	// This parameter is required.
	//
	// example:
	//
	// PERSONAL
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PERSONAL
	Accounts []*CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList) GetAccounts() []*CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts {
	return s.Accounts
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList) SetAccountType(v string) *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList {
	s.AccountType = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList) SetAccounts(v []*CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts) *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList {
	s.Accounts = v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingList) Validate() error {
	return dara.Validate(s)
}

type CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts struct {
	// This parameter is required.
	//
	// example:
	//
	// 300001111
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts) SetAccountId(v string) *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts {
	s.AccountId = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandRulesUserMappingListAccounts) Validate() error {
	return dara.Validate(s)
}

type CreateRowPermissionRequestCreateRowPermissionCommandTables struct {
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

func (s CreateRowPermissionRequestCreateRowPermissionCommandTables) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionRequestCreateRowPermissionCommandTables) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandTables) GetColumnName() *string {
	return s.ColumnName
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandTables) GetMappingColumnName() *string {
	return s.MappingColumnName
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandTables) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandTables) SetColumnName(v string) *CreateRowPermissionRequestCreateRowPermissionCommandTables {
	s.ColumnName = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandTables) SetMappingColumnName(v string) *CreateRowPermissionRequestCreateRowPermissionCommandTables {
	s.MappingColumnName = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandTables) SetResourceId(v string) *CreateRowPermissionRequestCreateRowPermissionCommandTables {
	s.ResourceId = &v
	return s
}

func (s *CreateRowPermissionRequestCreateRowPermissionCommandTables) Validate() error {
	return dara.Validate(s)
}
