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
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of the clusters that belong to your Alibaba Cloud account, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DefaultAlgo *string `json:"DefaultAlgo,omitempty" xml:"DefaultAlgo,omitempty"`
	// Specifies whether to enable the specified masking rule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > This parameter is valid only when the `RuleNameList` parameter is specfied.
	//
	// example:
	//
	// true
	Enable           *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	InterfaceVersion *string `json:"InterfaceVersion,omitempty" xml:"InterfaceVersion,omitempty"`
	MaskingAlgo      *string `json:"MaskingAlgo,omitempty" xml:"MaskingAlgo,omitempty"`
	// The parameter that is used to specify the masking rule that you want to modify and the value in the JSON format. All parameter values are of the string type. Example: `{"auto": {"databases": ["db1"], "tables": ["tb1"], "columns": ["c1,c2"] }, "description": "This rule will be applied to the columns c1 and c2 in table t1", "enabled": true, "applies_to": ["user"]}`. Where,
	//
	// 	- `"auto"`: specifies that the dynamic masking algorithm is supported. This parameter is required.
	//
	// 	- `"databases"`: Optional. The names of databases to which the masking rule is applied. Separate the names with commas (,). If you leave this parameter empty, the masking rule applies to all databases in the cluster.
	//
	// 	- `"tables"`: Optional. The names of tables to which the masking rule is applied. Separate the names with commas (,). If you leave this parameter empty, the rule applies to all tables in the cluster.
	//
	// 	- `"columns"`: Required. The names of fields to which the masking rule is applied. Separate the names with commas (,).
	//
	// 	- `"description"`: Optional. The description of the masking rule. The description is up to 64 characters in length.
	//
	// 	- `"enabled"`: Required. Specifies whether to enable the masking rule. Valid values: **true*	- (enable) and **false*	- (disable).
	//
	// 	- `"applies_to"`: The names of database accounts to which the masking rule is applied. Separate the names with commas (,).
	//
	// 	- `"exempted"`: The names of database accounts to which the masking rule is not applied. Separate the names with commas (,).
	//
	// >
	//
	// 	- If you specify `RuleName`, `RuleConfig` parameter is required.
	//
	// 	- You need to select either `"applies_to"` or `"exempted"`.
	//
	// example:
	//
	// {"auto": {"databases": ["db1"], "tables": ["tb1"], "columns": ["c1,c2"] }, "description": "This rule will be applied to the columns c1 and c2 in table t1", "enabled": true, "applies_to": ["user"]}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The name of the data masking rule. You can specify only one rule name at a time.
	//
	// >
	//
	// 	- You can call the [DescribeMaskingRules](https://help.aliyun.com/document_detail/212573.html) operation to query the details of all masking rules for a specified cluster, such as the names of the masking rules.
	//
	// 	- If the rule name does not exist in the cluster, the system automatically creates a masking rule based on the name and the value of `RuleConfig`.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The list of masking rule names. You can specify one or more masking rules at a time. Separate the masking rule names with commas (,).
	//
	// > You must specify either the `RuleName` or `RuleNameList` parameter.
	//
	// example:
	//
	// testrule
	RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
	// The version of the masking rule. Default value: v1. Valid values:
	//
	// 	- v1
	//
	// 	- v2
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
