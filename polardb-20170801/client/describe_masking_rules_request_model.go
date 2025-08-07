// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMaskingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeMaskingRulesRequest
	GetDBClusterId() *string
	SetInterfaceVersion(v string) *DescribeMaskingRulesRequest
	GetInterfaceVersion() *string
	SetRuleNameList(v string) *DescribeMaskingRulesRequest
	GetRuleNameList() *string
}

type DescribeMaskingRulesRequest struct {
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
	// Queries data masking rules or encryption rules. Valid values:
	//
	// v1: queries data masking rules. v2: queries data encryption rules.
	//
	// example:
	//
	// v1
	InterfaceVersion *string `json:"InterfaceVersion,omitempty" xml:"InterfaceVersion,omitempty"`
	// The name of the masking rule.
	//
	// example:
	//
	// testrule
	RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
}

func (s DescribeMaskingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeMaskingRulesRequest) GetInterfaceVersion() *string {
	return s.InterfaceVersion
}

func (s *DescribeMaskingRulesRequest) GetRuleNameList() *string {
	return s.RuleNameList
}

func (s *DescribeMaskingRulesRequest) SetDBClusterId(v string) *DescribeMaskingRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMaskingRulesRequest) SetInterfaceVersion(v string) *DescribeMaskingRulesRequest {
	s.InterfaceVersion = &v
	return s
}

func (s *DescribeMaskingRulesRequest) SetRuleNameList(v string) *DescribeMaskingRulesRequest {
	s.RuleNameList = &v
	return s
}

func (s *DescribeMaskingRulesRequest) Validate() error {
	return dara.Validate(s)
}
