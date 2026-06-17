// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaskingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyMaskingRulesRequest
	GetDBClusterId() *string
	SetDefaultAlgo(v string) *ModifyMaskingRulesRequest
	GetDefaultAlgo() *string
	SetEnable(v string) *ModifyMaskingRulesRequest
	GetEnable() *string
	SetInterfaceVersion(v string) *ModifyMaskingRulesRequest
	GetInterfaceVersion() *string
	SetMaskingAlgo(v string) *ModifyMaskingRulesRequest
	GetMaskingAlgo() *string
	SetRuleConfig(v string) *ModifyMaskingRulesRequest
	GetRuleConfig() *string
	SetRuleName(v string) *ModifyMaskingRulesRequest
	GetRuleName() *string
	SetRuleNameList(v string) *ModifyMaskingRulesRequest
	GetRuleNameList() *string
	SetRuleVersion(v string) *ModifyMaskingRulesRequest
	GetRuleVersion() *string
}

type ModifyMaskingRulesRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters in your account, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The default algorithm.
	//
	// > You must specify either MaskingAlgo or DefaultAIgo.
	//
	// example:
	//
	// aes-128-gcm
	DefaultAlgo *string `json:"DefaultAlgo,omitempty" xml:"DefaultAlgo,omitempty"`
	// Enables or disables the specified data masking rules. Valid values:
	//
	// - **true**: enables the specified rules.
	//
	// - **false**: disables the specified rules.
	//
	// > This parameter applies only when the `RuleNameList` parameter is specified.
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of rule to modify. Valid values:
	//
	// v1: Modifies a data masking rule.
	//
	// v2: Modifies an encryption rule.
	//
	// example:
	//
	// v1
	InterfaceVersion *string `json:"InterfaceVersion,omitempty" xml:"InterfaceVersion,omitempty"`
	// The masking algorithm. Specify one or more algorithms and their parameters. Format: `[{ "name": "algorithm_name", "params": {"key": "value"} }]`
	//
	// example:
	//
	// [{
	//
	// "name": "aes-128-gcm"
	//
	// }]
	MaskingAlgo *string `json:"MaskingAlgo,omitempty" xml:"MaskingAlgo,omitempty"`
	// A JSON string that specifies the rule configuration. Example: `{"auto": {"databases": ["db1"], "tables": ["tb1"], "columns": ["c1,c2"] }, "description": "This rule will be applied to the columns c1 and c2 in table t1", "enabled": true, "applies_to": ["user"]}`. The JSON string includes the following fields:
	//
	// - `"auto"`: Required. The object that contains the configuration for the dynamic data masking algorithm.
	//
	// - `"databases"`: Optional. The databases to which the rule applies. Separate multiple database names with a comma (,). If this parameter is omitted, the rule applies to all databases in the cluster.
	//
	// - `"tables"`: Optional. The tables to which the rule applies. Separate multiple table names with a comma (,). If this parameter is omitted, the rule applies to all tables in the cluster.
	//
	// - `"columns"`: Required. The columns to which the rule applies. Separate multiple column names with a comma (,).
	//
	// - `"description"`: Optional. The rule description, up to 64 characters in length.
	//
	// - `"enabled"`: Required. Specifies whether the data masking rule is enabled. Valid values: **true*	- (enabled) and **false*	- (disabled).
	//
	// - `"applies_to"`: The database accounts to which the rule applies. Separate multiple account names with a comma (,).
	//
	// - `"exempted"`: The database accounts that are exempt from the rule. Separate multiple account names with a comma (,).
	//
	// > 	- If you specify the `RuleName` parameter, you must also specify the `RuleConfig` parameter.
	//
	// >
	//
	// > 	- You must specify either `"applies_to"` or `"exempted"`.
	//
	// example:
	//
	// {"auto": {"databases": ["db1"], "tables": ["tb1"], "columns": ["c1,c2"] }, "description": "This rule will be applied to the columns c1 and c2 in table t1", "enabled": true, "applies_to": ["user"]}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The name of the data masking rule. You can specify only one rule name at a time.
	//
	// > - You can call the [DescribeMaskingRules](https://help.aliyun.com/document_detail/212573.html) operation to query the details of all data masking rules in the target cluster, including rule names.
	//
	// >
	//
	// > - If a rule with the specified name does not exist, the system creates a new one based on the provided `RuleConfig`.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// A comma-separated list of data masking rule names.
	//
	// > You must specify either the `RuleName` or `RuleNameList` parameter.
	//
	// example:
	//
	// testrule
	RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
	// The version of the data masking rule. Valid values:
	//
	// - v1 (default)
	//
	// - v2
	//
	// example:
	//
	// v2
	RuleVersion *string `json:"RuleVersion,omitempty" xml:"RuleVersion,omitempty"`
}

func (s ModifyMaskingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaskingRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaskingRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyMaskingRulesRequest) GetDefaultAlgo() *string {
	return s.DefaultAlgo
}

func (s *ModifyMaskingRulesRequest) GetEnable() *string {
	return s.Enable
}

func (s *ModifyMaskingRulesRequest) GetInterfaceVersion() *string {
	return s.InterfaceVersion
}

func (s *ModifyMaskingRulesRequest) GetMaskingAlgo() *string {
	return s.MaskingAlgo
}

func (s *ModifyMaskingRulesRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *ModifyMaskingRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyMaskingRulesRequest) GetRuleNameList() *string {
	return s.RuleNameList
}

func (s *ModifyMaskingRulesRequest) GetRuleVersion() *string {
	return s.RuleVersion
}

func (s *ModifyMaskingRulesRequest) SetDBClusterId(v string) *ModifyMaskingRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetDefaultAlgo(v string) *ModifyMaskingRulesRequest {
	s.DefaultAlgo = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetEnable(v string) *ModifyMaskingRulesRequest {
	s.Enable = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetInterfaceVersion(v string) *ModifyMaskingRulesRequest {
	s.InterfaceVersion = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetMaskingAlgo(v string) *ModifyMaskingRulesRequest {
	s.MaskingAlgo = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetRuleConfig(v string) *ModifyMaskingRulesRequest {
	s.RuleConfig = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetRuleName(v string) *ModifyMaskingRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetRuleNameList(v string) *ModifyMaskingRulesRequest {
	s.RuleNameList = &v
	return s
}

func (s *ModifyMaskingRulesRequest) SetRuleVersion(v string) *ModifyMaskingRulesRequest {
	s.RuleVersion = &v
	return s
}

func (s *ModifyMaskingRulesRequest) Validate() error {
	return dara.Validate(s)
}
