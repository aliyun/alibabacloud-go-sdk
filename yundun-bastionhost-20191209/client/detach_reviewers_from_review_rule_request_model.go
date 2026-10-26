// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachReviewersFromReviewRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DetachReviewersFromReviewRuleRequest
	GetInstanceId() *string
	SetPrincipalIds(v string) *DetachReviewersFromReviewRuleRequest
	GetPrincipalIds() *string
	SetRegionId(v string) *DetachReviewersFromReviewRuleRequest
	GetRegionId() *string
	SetReviewRuleId(v int64) *DetachReviewersFromReviewRuleRequest
	GetReviewRuleId() *int64
}

type DetachReviewersFromReviewRuleRequest struct {
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the auditors to remove.
	//
	// > You can call the GetReviewRule operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	PrincipalIds *string `json:"PrincipalIds,omitempty" xml:"PrincipalIds,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the audit policy.
	//
	// > You can invoke the ListReviewRules operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewRuleId *int64 `json:"ReviewRuleId,omitempty" xml:"ReviewRuleId,omitempty"`
}

func (s DetachReviewersFromReviewRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachReviewersFromReviewRuleRequest) GoString() string {
	return s.String()
}

func (s *DetachReviewersFromReviewRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachReviewersFromReviewRuleRequest) GetPrincipalIds() *string {
	return s.PrincipalIds
}

func (s *DetachReviewersFromReviewRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachReviewersFromReviewRuleRequest) GetReviewRuleId() *int64 {
	return s.ReviewRuleId
}

func (s *DetachReviewersFromReviewRuleRequest) SetInstanceId(v string) *DetachReviewersFromReviewRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachReviewersFromReviewRuleRequest) SetPrincipalIds(v string) *DetachReviewersFromReviewRuleRequest {
	s.PrincipalIds = &v
	return s
}

func (s *DetachReviewersFromReviewRuleRequest) SetRegionId(v string) *DetachReviewersFromReviewRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DetachReviewersFromReviewRuleRequest) SetReviewRuleId(v int64) *DetachReviewersFromReviewRuleRequest {
	s.ReviewRuleId = &v
	return s
}

func (s *DetachReviewersFromReviewRuleRequest) Validate() error {
	return dara.Validate(s)
}
