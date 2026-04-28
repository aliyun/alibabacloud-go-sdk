// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostOptimizationOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunId(v int64) *DescribeCostOptimizationOverviewRequest
	GetAssumeAliyunId() *int64
	SetAssumeAliyunIdList(v []*int64) *DescribeCostOptimizationOverviewRequest
	GetAssumeAliyunIdList() []*int64
	SetCheckPlanId(v int64) *DescribeCostOptimizationOverviewRequest
	GetCheckPlanId() *int64
	SetToken(v string) *DescribeCostOptimizationOverviewRequest
	GetToken() *string
}

type DescribeCostOptimizationOverviewRequest struct {
	// example:
	//
	// 11***********35
	AssumeAliyunId     *int64   `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	AssumeAliyunIdList []*int64 `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty" type:"Repeated"`
	CheckPlanId        *int64   `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeCostOptimizationOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostOptimizationOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewRequest) GetAssumeAliyunId() *int64 {
	return s.AssumeAliyunId
}

func (s *DescribeCostOptimizationOverviewRequest) GetAssumeAliyunIdList() []*int64 {
	return s.AssumeAliyunIdList
}

func (s *DescribeCostOptimizationOverviewRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeCostOptimizationOverviewRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeCostOptimizationOverviewRequest) SetAssumeAliyunId(v int64) *DescribeCostOptimizationOverviewRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewRequest) SetAssumeAliyunIdList(v []*int64) *DescribeCostOptimizationOverviewRequest {
	s.AssumeAliyunIdList = v
	return s
}

func (s *DescribeCostOptimizationOverviewRequest) SetCheckPlanId(v int64) *DescribeCostOptimizationOverviewRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewRequest) SetToken(v string) *DescribeCostOptimizationOverviewRequest {
	s.Token = &v
	return s
}

func (s *DescribeCostOptimizationOverviewRequest) Validate() error {
	return dara.Validate(s)
}
