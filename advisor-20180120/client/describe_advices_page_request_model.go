// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdviceId(v int64) *DescribeAdvicesPageRequest
	GetAdviceId() *int64
	SetCheckId(v string) *DescribeAdvicesPageRequest
	GetCheckId() *string
	SetCheckPlanId(v int64) *DescribeAdvicesPageRequest
	GetCheckPlanId() *int64
	SetLanguage(v string) *DescribeAdvicesPageRequest
	GetLanguage() *string
	SetPageNumber(v int32) *DescribeAdvicesPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAdvicesPageRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAdvicesPageRequest
	GetProduct() *string
	SetResourceId(v string) *DescribeAdvicesPageRequest
	GetResourceId() *string
}

type DescribeAdvicesPageRequest struct {
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
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvicesPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesPageRequest) GetAdviceId() *int64 {
	return s.AdviceId
}

func (s *DescribeAdvicesPageRequest) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeAdvicesPageRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeAdvicesPageRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAdvicesPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAdvicesPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAdvicesPageRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvicesPageRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAdvicesPageRequest) SetAdviceId(v int64) *DescribeAdvicesPageRequest {
	s.AdviceId = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetCheckId(v string) *DescribeAdvicesPageRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetCheckPlanId(v int64) *DescribeAdvicesPageRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetLanguage(v string) *DescribeAdvicesPageRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetPageNumber(v int32) *DescribeAdvicesPageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetPageSize(v int32) *DescribeAdvicesPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetProduct(v string) *DescribeAdvicesPageRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesPageRequest) SetResourceId(v string) *DescribeAdvicesPageRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesPageRequest) Validate() error {
	return dara.Validate(s)
}
