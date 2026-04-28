// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorCostCheckShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunIdListShrink(v string) *RefreshAdvisorCostCheckShrinkRequest
	GetAssumeAliyunIdListShrink() *string
	SetCheckIdsShrink(v string) *RefreshAdvisorCostCheckShrinkRequest
	GetCheckIdsShrink() *string
	SetCheckPlanId(v int64) *RefreshAdvisorCostCheckShrinkRequest
	GetCheckPlanId() *int64
	SetProduct(v string) *RefreshAdvisorCostCheckShrinkRequest
	GetProduct() *string
	SetRefreshResource(v bool) *RefreshAdvisorCostCheckShrinkRequest
	GetRefreshResource() *bool
	SetResourceIdsShrink(v string) *RefreshAdvisorCostCheckShrinkRequest
	GetResourceIdsShrink() *string
}

type RefreshAdvisorCostCheckShrinkRequest struct {
	AssumeAliyunIdListShrink *string `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty"`
	CheckIdsShrink           *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	CheckPlanId              *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// false
	RefreshResource   *bool   `json:"RefreshResource,omitempty" xml:"RefreshResource,omitempty"`
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
}

func (s RefreshAdvisorCostCheckShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCostCheckShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckShrinkRequest) GetAssumeAliyunIdListShrink() *string {
	return s.AssumeAliyunIdListShrink
}

func (s *RefreshAdvisorCostCheckShrinkRequest) GetCheckIdsShrink() *string {
	return s.CheckIdsShrink
}

func (s *RefreshAdvisorCostCheckShrinkRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *RefreshAdvisorCostCheckShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *RefreshAdvisorCostCheckShrinkRequest) GetRefreshResource() *bool {
	return s.RefreshResource
}

func (s *RefreshAdvisorCostCheckShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetAssumeAliyunIdListShrink(v string) *RefreshAdvisorCostCheckShrinkRequest {
	s.AssumeAliyunIdListShrink = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetCheckIdsShrink(v string) *RefreshAdvisorCostCheckShrinkRequest {
	s.CheckIdsShrink = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetCheckPlanId(v int64) *RefreshAdvisorCostCheckShrinkRequest {
	s.CheckPlanId = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetProduct(v string) *RefreshAdvisorCostCheckShrinkRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetRefreshResource(v bool) *RefreshAdvisorCostCheckShrinkRequest {
	s.RefreshResource = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) SetResourceIdsShrink(v string) *RefreshAdvisorCostCheckShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *RefreshAdvisorCostCheckShrinkRequest) Validate() error {
	return dara.Validate(s)
}
