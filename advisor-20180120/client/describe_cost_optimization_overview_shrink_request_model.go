// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostOptimizationOverviewShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunId(v int64) *DescribeCostOptimizationOverviewShrinkRequest
	GetAssumeAliyunId() *int64
	SetAssumeAliyunIdListShrink(v string) *DescribeCostOptimizationOverviewShrinkRequest
	GetAssumeAliyunIdListShrink() *string
	SetCheckPlanId(v int64) *DescribeCostOptimizationOverviewShrinkRequest
	GetCheckPlanId() *int64
	SetToken(v string) *DescribeCostOptimizationOverviewShrinkRequest
	GetToken() *string
}

type DescribeCostOptimizationOverviewShrinkRequest struct {
	// example:
	//
	// 11***********35
	AssumeAliyunId           *int64  `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	AssumeAliyunIdListShrink *string `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty"`
	CheckPlanId              *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeCostOptimizationOverviewShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostOptimizationOverviewShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) GetAssumeAliyunId() *int64 {
	return s.AssumeAliyunId
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) GetAssumeAliyunIdListShrink() *string {
	return s.AssumeAliyunIdListShrink
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) SetAssumeAliyunId(v int64) *DescribeCostOptimizationOverviewShrinkRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) SetAssumeAliyunIdListShrink(v string) *DescribeCostOptimizationOverviewShrinkRequest {
	s.AssumeAliyunIdListShrink = &v
	return s
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) SetCheckPlanId(v int64) *DescribeCostOptimizationOverviewShrinkRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) SetToken(v string) *DescribeCostOptimizationOverviewShrinkRequest {
	s.Token = &v
	return s
}

func (s *DescribeCostOptimizationOverviewShrinkRequest) Validate() error {
	return dara.Validate(s)
}
