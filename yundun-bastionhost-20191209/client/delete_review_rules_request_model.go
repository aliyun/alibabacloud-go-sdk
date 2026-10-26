// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReviewRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteReviewRulesRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteReviewRulesRequest
	GetRegionId() *string
	SetReviewRuleIds(v string) *DeleteReviewRulesRequest
	GetReviewRuleIds() *string
}

type DeleteReviewRulesRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The array of audit policy IDs to delete.
	//
	// > You can call the ListReviewRules operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2,3
	ReviewRuleIds *string `json:"ReviewRuleIds,omitempty" xml:"ReviewRuleIds,omitempty"`
}

func (s DeleteReviewRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReviewRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteReviewRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteReviewRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteReviewRulesRequest) GetReviewRuleIds() *string {
	return s.ReviewRuleIds
}

func (s *DeleteReviewRulesRequest) SetInstanceId(v string) *DeleteReviewRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteReviewRulesRequest) SetRegionId(v string) *DeleteReviewRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteReviewRulesRequest) SetReviewRuleIds(v string) *DeleteReviewRulesRequest {
	s.ReviewRuleIds = &v
	return s
}

func (s *DeleteReviewRulesRequest) Validate() error {
	return dara.Validate(s)
}
