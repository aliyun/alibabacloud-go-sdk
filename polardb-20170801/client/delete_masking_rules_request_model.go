// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaskingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteMaskingRulesRequest
	GetDBClusterId() *string
	SetInterfaceVersion(v string) *DeleteMaskingRulesRequest
	GetInterfaceVersion() *string
	SetRuleNameList(v string) *DeleteMaskingRulesRequest
	GetRuleNameList() *string
}

type DeleteMaskingRulesRequest struct {
	// The cluster ID.
	//
	// > For more information, see [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The type of rule to delete. Valid values:
	//
	// v1: deletes data masking rules.
	//
	// v2: deletes data encryption rules.
	//
	// example:
	//
	// v1
	InterfaceVersion *string `json:"InterfaceVersion,omitempty" xml:"InterfaceVersion,omitempty"`
	// The names of the data masking rules to delete. To delete multiple rules in a batch, separate the names with commas (,).
	//
	// > For more information, see [DescribeMaskingRules](https://help.aliyun.com/document_detail/212573.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
}

func (s DeleteMaskingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaskingRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaskingRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteMaskingRulesRequest) GetInterfaceVersion() *string {
	return s.InterfaceVersion
}

func (s *DeleteMaskingRulesRequest) GetRuleNameList() *string {
	return s.RuleNameList
}

func (s *DeleteMaskingRulesRequest) SetDBClusterId(v string) *DeleteMaskingRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteMaskingRulesRequest) SetInterfaceVersion(v string) *DeleteMaskingRulesRequest {
	s.InterfaceVersion = &v
	return s
}

func (s *DeleteMaskingRulesRequest) SetRuleNameList(v string) *DeleteMaskingRulesRequest {
	s.RuleNameList = &v
	return s
}

func (s *DeleteMaskingRulesRequest) Validate() error {
	return dara.Validate(s)
}
