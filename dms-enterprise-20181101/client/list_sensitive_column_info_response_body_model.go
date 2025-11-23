// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSensitiveColumnInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSensitiveColumnInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSensitiveColumnInfoResponseBody
	GetRequestId() *string
	SetSensitiveColumnList(v *ListSensitiveColumnInfoResponseBodySensitiveColumnList) *ListSensitiveColumnInfoResponseBody
	GetSensitiveColumnList() *ListSensitiveColumnInfoResponseBodySensitiveColumnList
	SetSuccess(v bool) *ListSensitiveColumnInfoResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSensitiveColumnInfoResponseBody
	GetTotalCount() *int64
}

type ListSensitiveColumnInfoResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the sensitive field.
	SensitiveColumnList *ListSensitiveColumnInfoResponseBodySensitiveColumnList `json:"SensitiveColumnList,omitempty" xml:"SensitiveColumnList,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSensitiveColumnInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSensitiveColumnInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSensitiveColumnInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSensitiveColumnInfoResponseBody) GetSensitiveColumnList() *ListSensitiveColumnInfoResponseBodySensitiveColumnList {
	return s.SensitiveColumnList
}

func (s *ListSensitiveColumnInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSensitiveColumnInfoResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSensitiveColumnInfoResponseBody) SetErrorCode(v string) *ListSensitiveColumnInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBody) SetErrorMessage(v string) *ListSensitiveColumnInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBody) SetRequestId(v string) *ListSensitiveColumnInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBody) SetSensitiveColumnList(v *ListSensitiveColumnInfoResponseBodySensitiveColumnList) *ListSensitiveColumnInfoResponseBody {
	s.SensitiveColumnList = v
	return s
}

func (s *ListSensitiveColumnInfoResponseBody) SetSuccess(v bool) *ListSensitiveColumnInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBody) SetTotalCount(v int64) *ListSensitiveColumnInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBody) Validate() error {
	if s.SensitiveColumnList != nil {
		if err := s.SensitiveColumnList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSensitiveColumnInfoResponseBodySensitiveColumnList struct {
	SensitiveColumn []*ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn `json:"SensitiveColumn,omitempty" xml:"SensitiveColumn,omitempty" type:"Repeated"`
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnList) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnList) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnList) GetSensitiveColumn() []*ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	return s.SensitiveColumn
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnList) SetSensitiveColumn(v []*ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) *ListSensitiveColumnInfoResponseBodySensitiveColumnList {
	s.SensitiveColumn = v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnList) Validate() error {
	if s.SensitiveColumn != nil {
		for _, item := range s.SensitiveColumn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn struct {
	// The name of the category.
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The name of the sensitive field.
	//
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The information about the default masking algorithm.
	DefaultDesensitizationRule *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule `json:"DefaultDesensitizationRule,omitempty" xml:"DefaultDesensitizationRule,omitempty" type:"Struct"`
	// The ID of the instance.
	//
	// example:
	//
	// 183****
	InstanceId *int32 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the sensitive field is displayed in plaintext.
	//
	// example:
	//
	// False
	IsPlain *bool `json:"IsPlain,omitempty" xml:"IsPlain,omitempty"`
	// The sample data.
	SampleData *string `json:"SampleData,omitempty" xml:"SampleData,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The sensitivity level of the field. Valid values:
	//
	// 	- Low
	//
	// 	- Medium
	//
	// 	- High
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The list of partial masking algorithms.
	SemiDesensitizationRuleList *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList `json:"SemiDesensitizationRuleList,omitempty" xml:"SemiDesensitizationRuleList,omitempty" type:"Struct"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The user-defined sensitivity level.
	//
	// example:
	//
	// S1
	UserSensitivityLevel *string `json:"UserSensitivityLevel,omitempty" xml:"UserSensitivityLevel,omitempty"`
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetDefaultDesensitizationRule() *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule {
	return s.DefaultDesensitizationRule
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetIsPlain() *bool {
	return s.IsPlain
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetSampleData() *string {
	return s.SampleData
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetSemiDesensitizationRuleList() *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList {
	return s.SemiDesensitizationRuleList
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) GetUserSensitivityLevel() *string {
	return s.UserSensitivityLevel
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetCategoryName(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.CategoryName = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetColumnName(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetDefaultDesensitizationRule(v *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.DefaultDesensitizationRule = v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetInstanceId(v int32) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.InstanceId = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetIsPlain(v bool) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.IsPlain = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetSampleData(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.SampleData = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetSchemaName(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetSecurityLevel(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.SecurityLevel = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetSemiDesensitizationRuleList(v *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.SemiDesensitizationRuleList = v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetTableName(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) SetUserSensitivityLevel(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn {
	s.UserSensitivityLevel = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumn) Validate() error {
	if s.DefaultDesensitizationRule != nil {
		if err := s.DefaultDesensitizationRule.Validate(); err != nil {
			return err
		}
	}
	if s.SemiDesensitizationRuleList != nil {
		if err := s.SemiDesensitizationRuleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule struct {
	// The masking algorithm ID.
	//
	// example:
	//
	// 101**
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The masking algorithm name.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule) SetRuleId(v int64) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule {
	s.RuleId = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule) SetRuleName(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule {
	s.RuleName = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnDefaultDesensitizationRule) Validate() error {
	return dara.Validate(s)
}

type ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList struct {
	SemiDesensitizationRule []*ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule `json:"SemiDesensitizationRule,omitempty" xml:"SemiDesensitizationRule,omitempty" type:"Repeated"`
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList) GetSemiDesensitizationRule() []*ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule {
	return s.SemiDesensitizationRule
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList) SetSemiDesensitizationRule(v []*ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList {
	s.SemiDesensitizationRule = v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleList) Validate() error {
	if s.SemiDesensitizationRule != nil {
		for _, item := range s.SemiDesensitizationRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule struct {
	// The ID of the partial masking algorithm.
	//
	// example:
	//
	// 10***
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The partial masking algorithm name.
	//
	// example:
	//
	// test01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule) SetRuleId(v int64) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule {
	s.RuleId = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule) SetRuleName(v string) *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule {
	s.RuleName = &v
	return s
}

func (s *ListSensitiveColumnInfoResponseBodySensitiveColumnListSensitiveColumnSemiDesensitizationRuleListSemiDesensitizationRule) Validate() error {
	return dara.Validate(s)
}
