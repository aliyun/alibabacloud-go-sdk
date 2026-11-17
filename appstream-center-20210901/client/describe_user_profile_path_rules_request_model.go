// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserProfilePathRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *DescribeUserProfilePathRulesRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupSetId(v string) *DescribeUserProfilePathRulesRequest
	GetAppInstanceGroupSetId() *string
	SetRegionId(v string) *DescribeUserProfilePathRulesRequest
	GetRegionId() *string
	SetRuleType(v string) *DescribeUserProfilePathRulesRequest
	GetRuleType() *string
}

type DescribeUserProfilePathRulesRequest struct {
	// The app instance group ID. Specify either this parameter or AppInstanceGroupSetId when querying app instance group rules. This parameter can be omitted when querying system default rules.
	//
	// example:
	//
	// aig-0bi83v****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The app instance group set ID. Specify either this parameter or AppInstanceGroupId when querying app instance group set rules. This parameter can be omitted when querying system default rules.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The region ID of the request. When querying or setting rules for a specified app instance group, use the region where the target app instance group is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The path rule type. Valid values:
	//
	// - `Default`: Queries system default rules. Both target IDs can be omitted.
	//
	// - `AppInstanceGroup`: Queries the rules of an app instance group or an app instance group set.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroup
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeUserProfilePathRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DescribeUserProfilePathRulesRequest) GetAppInstanceGroupSetId() *string {
	return s.AppInstanceGroupSetId
}

func (s *DescribeUserProfilePathRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserProfilePathRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeUserProfilePathRulesRequest) SetAppInstanceGroupId(v string) *DescribeUserProfilePathRulesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeUserProfilePathRulesRequest) SetAppInstanceGroupSetId(v string) *DescribeUserProfilePathRulesRequest {
	s.AppInstanceGroupSetId = &v
	return s
}

func (s *DescribeUserProfilePathRulesRequest) SetRegionId(v string) *DescribeUserProfilePathRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserProfilePathRulesRequest) SetRuleType(v string) *DescribeUserProfilePathRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeUserProfilePathRulesRequest) Validate() error {
	return dara.Validate(s)
}
