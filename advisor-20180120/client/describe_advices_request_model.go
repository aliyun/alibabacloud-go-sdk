// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdviceId(v int64) *DescribeAdvicesRequest
	GetAdviceId() *int64
	SetCheckId(v string) *DescribeAdvicesRequest
	GetCheckId() *string
	SetCheckPlanId(v int64) *DescribeAdvicesRequest
	GetCheckPlanId() *int64
	SetExcludeAdviceId(v int64) *DescribeAdvicesRequest
	GetExcludeAdviceId() *int64
	SetLanguage(v string) *DescribeAdvicesRequest
	GetLanguage() *string
	SetProduct(v string) *DescribeAdvicesRequest
	GetProduct() *string
	SetResourceId(v string) *DescribeAdvicesRequest
	GetResourceId() *string
}

type DescribeAdvicesRequest struct {
	// example:
	//
	// 12345678
	AdviceId *int64 `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// 12345678
	ExcludeAdviceId *int64 `json:"ExcludeAdviceId,omitempty" xml:"ExcludeAdviceId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesRequest) GetAdviceId() *int64 {
	return s.AdviceId
}

func (s *DescribeAdvicesRequest) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeAdvicesRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeAdvicesRequest) GetExcludeAdviceId() *int64 {
	return s.ExcludeAdviceId
}

func (s *DescribeAdvicesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAdvicesRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvicesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAdvicesRequest) SetAdviceId(v int64) *DescribeAdvicesRequest {
	s.AdviceId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetCheckId(v string) *DescribeAdvicesRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetCheckPlanId(v int64) *DescribeAdvicesRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetExcludeAdviceId(v int64) *DescribeAdvicesRequest {
	s.ExcludeAdviceId = &v
	return s
}

func (s *DescribeAdvicesRequest) SetLanguage(v string) *DescribeAdvicesRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvicesRequest) SetProduct(v string) *DescribeAdvicesRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesRequest) SetResourceId(v string) *DescribeAdvicesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesRequest) Validate() error {
	return dara.Validate(s)
}
