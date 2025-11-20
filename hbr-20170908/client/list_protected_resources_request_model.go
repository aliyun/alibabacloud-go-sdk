// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectedResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedByProduct(v string) *ListProtectedResourcesRequest
	GetCreatedByProduct() *string
	SetMaxResults(v int32) *ListProtectedResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProtectedResourcesRequest
	GetNextToken() *string
	SetResourceId(v string) *ListProtectedResourcesRequest
	GetResourceId() *string
	SetSkip(v int32) *ListProtectedResourcesRequest
	GetSkip() *int32
	SetSourceType(v string) *ListProtectedResourcesRequest
	GetSourceType() *string
}

type ListProtectedResourcesRequest struct {
	// example:
	//
	// BASIC
	CreatedByProduct *string `json:"CreatedByProduct,omitempty" xml:"CreatedByProduct,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// aWQj********MCMy
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// i-wz95************7zrd
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 10
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListProtectedResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListProtectedResourcesRequest) GetCreatedByProduct() *string {
	return s.CreatedByProduct
}

func (s *ListProtectedResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProtectedResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProtectedResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListProtectedResourcesRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProtectedResourcesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListProtectedResourcesRequest) SetCreatedByProduct(v string) *ListProtectedResourcesRequest {
	s.CreatedByProduct = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetMaxResults(v int32) *ListProtectedResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetNextToken(v string) *ListProtectedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetResourceId(v string) *ListProtectedResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetSkip(v int32) *ListProtectedResourcesRequest {
	s.Skip = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetSourceType(v string) *ListProtectedResourcesRequest {
	s.SourceType = &v
	return s
}

func (s *ListProtectedResourcesRequest) Validate() error {
	return dara.Validate(s)
}
