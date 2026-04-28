// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesFlatPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdviceId(v int64) *DescribeAdvicesFlatPageRequest
	GetAdviceId() *int64
	SetCheckId(v string) *DescribeAdvicesFlatPageRequest
	GetCheckId() *string
	SetLanguage(v string) *DescribeAdvicesFlatPageRequest
	GetLanguage() *string
	SetPageNumber(v int32) *DescribeAdvicesFlatPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAdvicesFlatPageRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAdvicesFlatPageRequest
	GetProduct() *string
	SetResourceId(v string) *DescribeAdvicesFlatPageRequest
	GetResourceId() *string
}

type DescribeAdvicesFlatPageRequest struct {
	// example:
	//
	// 12345678
	AdviceId *int64 `json:"AdviceId,omitempty" xml:"AdviceId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
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
	// i-2zecnwitr2s7aca6****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeAdvicesFlatPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesFlatPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesFlatPageRequest) GetAdviceId() *int64 {
	return s.AdviceId
}

func (s *DescribeAdvicesFlatPageRequest) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeAdvicesFlatPageRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAdvicesFlatPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAdvicesFlatPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAdvicesFlatPageRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvicesFlatPageRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAdvicesFlatPageRequest) SetAdviceId(v int64) *DescribeAdvicesFlatPageRequest {
	s.AdviceId = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetCheckId(v string) *DescribeAdvicesFlatPageRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetLanguage(v string) *DescribeAdvicesFlatPageRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetPageNumber(v int32) *DescribeAdvicesFlatPageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetPageSize(v int32) *DescribeAdvicesFlatPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetProduct(v string) *DescribeAdvicesFlatPageRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) SetResourceId(v string) *DescribeAdvicesFlatPageRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvicesFlatPageRequest) Validate() error {
	return dara.Validate(s)
}
