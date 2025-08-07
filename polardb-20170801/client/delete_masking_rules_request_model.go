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
	// Deletes data masking or encryption rules. Valid values:
	//
	// v1: deletes data masking rules. v2: deletes data encryption rules.
	//
	// example:
	//
	// v1
	InterfaceVersion *string `json:"InterfaceVersion,omitempty" xml:"InterfaceVersion,omitempty"`
	// The name of the masking rule. You can specify multiple masking rules at a time. Separate the masking rules with commas (,).
	//
	// > You can call the [DescribeMaskingRules](https://help.aliyun.com/document_detail/212573.html) operation to query details of all the masking rules for a specified cluster, such as the names of the masking rules.
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
