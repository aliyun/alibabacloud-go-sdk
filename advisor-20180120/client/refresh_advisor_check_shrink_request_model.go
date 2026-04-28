// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorCheckShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunId(v int64) *RefreshAdvisorCheckShrinkRequest
	GetAssumeAliyunId() *int64
	SetCheckId(v string) *RefreshAdvisorCheckShrinkRequest
	GetCheckId() *string
	SetCheckPlanId(v int64) *RefreshAdvisorCheckShrinkRequest
	GetCheckPlanId() *int64
	SetLanguage(v string) *RefreshAdvisorCheckShrinkRequest
	GetLanguage() *string
	SetProduct(v string) *RefreshAdvisorCheckShrinkRequest
	GetProduct() *string
	SetResourceDimensionListShrink(v string) *RefreshAdvisorCheckShrinkRequest
	GetResourceDimensionListShrink() *string
	SetResourceId(v string) *RefreshAdvisorCheckShrinkRequest
	GetResourceId() *string
	SetToken(v string) *RefreshAdvisorCheckShrinkRequest
	GetToken() *string
}

type RefreshAdvisorCheckShrinkRequest struct {
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product                     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceDimensionListShrink *string `json:"ResourceDimensionList,omitempty" xml:"ResourceDimensionList,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RefreshAdvisorCheckShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCheckShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckShrinkRequest) GetAssumeAliyunId() *int64 {
	return s.AssumeAliyunId
}

func (s *RefreshAdvisorCheckShrinkRequest) GetCheckId() *string {
	return s.CheckId
}

func (s *RefreshAdvisorCheckShrinkRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *RefreshAdvisorCheckShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *RefreshAdvisorCheckShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *RefreshAdvisorCheckShrinkRequest) GetResourceDimensionListShrink() *string {
	return s.ResourceDimensionListShrink
}

func (s *RefreshAdvisorCheckShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *RefreshAdvisorCheckShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *RefreshAdvisorCheckShrinkRequest) SetAssumeAliyunId(v int64) *RefreshAdvisorCheckShrinkRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetCheckId(v string) *RefreshAdvisorCheckShrinkRequest {
	s.CheckId = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetCheckPlanId(v int64) *RefreshAdvisorCheckShrinkRequest {
	s.CheckPlanId = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetLanguage(v string) *RefreshAdvisorCheckShrinkRequest {
	s.Language = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetProduct(v string) *RefreshAdvisorCheckShrinkRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetResourceDimensionListShrink(v string) *RefreshAdvisorCheckShrinkRequest {
	s.ResourceDimensionListShrink = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetResourceId(v string) *RefreshAdvisorCheckShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) SetToken(v string) *RefreshAdvisorCheckShrinkRequest {
	s.Token = &v
	return s
}

func (s *RefreshAdvisorCheckShrinkRequest) Validate() error {
	return dara.Validate(s)
}
