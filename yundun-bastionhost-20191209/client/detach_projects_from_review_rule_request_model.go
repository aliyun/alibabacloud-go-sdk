// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachProjectsFromReviewRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DetachProjectsFromReviewRuleRequest
	GetInstanceId() *string
	SetProjectIds(v string) *DetachProjectsFromReviewRuleRequest
	GetProjectIds() *string
	SetRegionId(v string) *DetachProjectsFromReviewRuleRequest
	GetRegionId() *string
	SetReviewRuleId(v int64) *DetachProjectsFromReviewRuleRequest
	GetReviewRuleId() *int64
}

type DetachProjectsFromReviewRuleRequest struct {
	// The instance ID of the bastion host.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The array of project IDs to remove from the audit policy.
	//
	// > You can call the GetReviewRule operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The audit policy ID.
	//
	// > You can call the ListReviewRules operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewRuleId *int64 `json:"ReviewRuleId,omitempty" xml:"ReviewRuleId,omitempty"`
}

func (s DetachProjectsFromReviewRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachProjectsFromReviewRuleRequest) GoString() string {
	return s.String()
}

func (s *DetachProjectsFromReviewRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachProjectsFromReviewRuleRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *DetachProjectsFromReviewRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachProjectsFromReviewRuleRequest) GetReviewRuleId() *int64 {
	return s.ReviewRuleId
}

func (s *DetachProjectsFromReviewRuleRequest) SetInstanceId(v string) *DetachProjectsFromReviewRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachProjectsFromReviewRuleRequest) SetProjectIds(v string) *DetachProjectsFromReviewRuleRequest {
	s.ProjectIds = &v
	return s
}

func (s *DetachProjectsFromReviewRuleRequest) SetRegionId(v string) *DetachProjectsFromReviewRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DetachProjectsFromReviewRuleRequest) SetReviewRuleId(v int64) *DetachProjectsFromReviewRuleRequest {
	s.ReviewRuleId = &v
	return s
}

func (s *DetachProjectsFromReviewRuleRequest) Validate() error {
	return dara.Validate(s)
}
