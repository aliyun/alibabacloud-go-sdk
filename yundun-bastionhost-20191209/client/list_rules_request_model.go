// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRulesRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListRulesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListRulesRequest
	GetPageSize() *string
	SetRegionId(v string) *ListRulesRequest
	GetRegionId() *string
	SetRuleName(v string) *ListRulesRequest
	GetRuleName() *string
	SetRuleState(v string) *ListRulesRequest
	GetRuleState() *string
}

type ListRulesRequest struct {
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-5yd34ol020a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Valid values: 1 to 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the authorization rule to query. Only exact match is supported.
	//
	// example:
	//
	// rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The state of the authorization rule to query.
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	RuleState *string `json:"RuleState,omitempty" xml:"RuleState,omitempty"`
}

func (s ListRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRulesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListRulesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRulesRequest) GetRuleState() *string {
	return s.RuleState
}

func (s *ListRulesRequest) SetInstanceId(v string) *ListRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRulesRequest) SetPageNumber(v string) *ListRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRulesRequest) SetPageSize(v string) *ListRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRulesRequest) SetRegionId(v string) *ListRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListRulesRequest) SetRuleName(v string) *ListRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListRulesRequest) SetRuleState(v string) *ListRulesRequest {
	s.RuleState = &v
	return s
}

func (s *ListRulesRequest) Validate() error {
	return dara.Validate(s)
}
