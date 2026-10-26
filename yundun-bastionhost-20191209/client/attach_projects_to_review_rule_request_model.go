// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachProjectsToReviewRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AttachProjectsToReviewRuleRequest
	GetInstanceId() *string
	SetProjectIds(v string) *AttachProjectsToReviewRuleRequest
	GetProjectIds() *string
	SetRegionId(v string) *AttachProjectsToReviewRuleRequest
	GetRegionId() *string
	SetReviewRuleId(v int64) *AttachProjectsToReviewRuleRequest
	GetReviewRuleId() *int64
}

type AttachProjectsToReviewRuleRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1ghxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the projects to associate.
	//
	// > You can call the ListProjects operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the audit policy.
	//
	// > You can invoke the ListReviewRules operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewRuleId *int64 `json:"ReviewRuleId,omitempty" xml:"ReviewRuleId,omitempty"`
}

func (s AttachProjectsToReviewRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachProjectsToReviewRuleRequest) GoString() string {
	return s.String()
}

func (s *AttachProjectsToReviewRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachProjectsToReviewRuleRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *AttachProjectsToReviewRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachProjectsToReviewRuleRequest) GetReviewRuleId() *int64 {
	return s.ReviewRuleId
}

func (s *AttachProjectsToReviewRuleRequest) SetInstanceId(v string) *AttachProjectsToReviewRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachProjectsToReviewRuleRequest) SetProjectIds(v string) *AttachProjectsToReviewRuleRequest {
	s.ProjectIds = &v
	return s
}

func (s *AttachProjectsToReviewRuleRequest) SetRegionId(v string) *AttachProjectsToReviewRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AttachProjectsToReviewRuleRequest) SetReviewRuleId(v int64) *AttachProjectsToReviewRuleRequest {
	s.ReviewRuleId = &v
	return s
}

func (s *AttachProjectsToReviewRuleRequest) Validate() error {
	return dara.Validate(s)
}
