// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSQLRateLimitingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteSQLRateLimitingRulesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteSQLRateLimitingRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSQLRateLimitingRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteSQLRateLimitingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSQLRateLimitingRulesRequest
	GetResourceOwnerId() *int64
	SetRuleNameList(v string) *DeleteSQLRateLimitingRulesRequest
	GetRuleNameList() *string
}

type DeleteSQLRateLimitingRulesRequest struct {
	// The cluster ID. Invocation instructions: append DPI engine type restrictions to the DBClusterId parameter description, specifying that the SQL throttling feature is supported only for PostgreSQL and Oracle DPI engine clusters, and MySQL clusters do not support this feature. Invocation instructions: append cluster status prerequisites to the DBClusterId parameter description, specifying that the cluster must be in the Running state when you invoke this API, and that you can invoke the DescribeDBClusters API and perform a Field-specific search on the DBClusterStatus field to query the current state.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The names of the SQL throttling rules to delete. You can specify multiple rule names to delete rules in batches. Separate multiple names with commas (,).
	//
	// > You can call the [DescribeSQLRateLimitingRules](https://help.aliyun.com/document_detail/212573.html) operation to view the details of all SQL throttling rules in the specified cluster, including the rule names.
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
}

func (s DeleteSQLRateLimitingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSQLRateLimitingRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteSQLRateLimitingRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteSQLRateLimitingRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSQLRateLimitingRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSQLRateLimitingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSQLRateLimitingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSQLRateLimitingRulesRequest) GetRuleNameList() *string {
	return s.RuleNameList
}

func (s *DeleteSQLRateLimitingRulesRequest) SetDBClusterId(v string) *DeleteSQLRateLimitingRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesRequest) SetOwnerAccount(v string) *DeleteSQLRateLimitingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesRequest) SetOwnerId(v int64) *DeleteSQLRateLimitingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesRequest) SetResourceOwnerAccount(v string) *DeleteSQLRateLimitingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesRequest) SetResourceOwnerId(v int64) *DeleteSQLRateLimitingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesRequest) SetRuleNameList(v string) *DeleteSQLRateLimitingRulesRequest {
	s.RuleNameList = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesRequest) Validate() error {
	return dara.Validate(s)
}
