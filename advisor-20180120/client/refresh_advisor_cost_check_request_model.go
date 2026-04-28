// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorCostCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunIdList(v []*int64) *RefreshAdvisorCostCheckRequest
	GetAssumeAliyunIdList() []*int64
	SetCheckIds(v []*string) *RefreshAdvisorCostCheckRequest
	GetCheckIds() []*string
	SetCheckPlanId(v int64) *RefreshAdvisorCostCheckRequest
	GetCheckPlanId() *int64
	SetProduct(v string) *RefreshAdvisorCostCheckRequest
	GetProduct() *string
	SetRefreshResource(v bool) *RefreshAdvisorCostCheckRequest
	GetRefreshResource() *bool
	SetResourceIds(v []*string) *RefreshAdvisorCostCheckRequest
	GetResourceIds() []*string
}

type RefreshAdvisorCostCheckRequest struct {
	AssumeAliyunIdList []*int64  `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty" type:"Repeated"`
	CheckIds           []*string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	CheckPlanId        *int64    `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// false
	RefreshResource *bool     `json:"RefreshResource,omitempty" xml:"RefreshResource,omitempty"`
	ResourceIds     []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s RefreshAdvisorCostCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCostCheckRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckRequest) GetAssumeAliyunIdList() []*int64 {
	return s.AssumeAliyunIdList
}

func (s *RefreshAdvisorCostCheckRequest) GetCheckIds() []*string {
	return s.CheckIds
}

func (s *RefreshAdvisorCostCheckRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *RefreshAdvisorCostCheckRequest) GetProduct() *string {
	return s.Product
}

func (s *RefreshAdvisorCostCheckRequest) GetRefreshResource() *bool {
	return s.RefreshResource
}

func (s *RefreshAdvisorCostCheckRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *RefreshAdvisorCostCheckRequest) SetAssumeAliyunIdList(v []*int64) *RefreshAdvisorCostCheckRequest {
	s.AssumeAliyunIdList = v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetCheckIds(v []*string) *RefreshAdvisorCostCheckRequest {
	s.CheckIds = v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetCheckPlanId(v int64) *RefreshAdvisorCostCheckRequest {
	s.CheckPlanId = &v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetProduct(v string) *RefreshAdvisorCostCheckRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetRefreshResource(v bool) *RefreshAdvisorCostCheckRequest {
	s.RefreshResource = &v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) SetResourceIds(v []*string) *RefreshAdvisorCostCheckRequest {
	s.ResourceIds = v
	return s
}

func (s *RefreshAdvisorCostCheckRequest) Validate() error {
	return dara.Validate(s)
}
