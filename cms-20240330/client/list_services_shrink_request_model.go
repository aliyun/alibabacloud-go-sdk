// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServicesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesShrinkRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListServicesShrinkRequest
	GetResourceGroupId() *string
	SetServiceName(v string) *ListServicesShrinkRequest
	GetServiceName() *string
	SetServiceType(v string) *ListServicesShrinkRequest
	GetServiceType() *string
	SetTagsShrink(v string) *ListServicesShrinkRequest
	GetTagsShrink() *string
}

type ListServicesShrinkRequest struct {
	// The maximum number of records to return in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Token for the next query, an empty value indicates the last page.
	//
	// example:
	//
	// 7-b81a-4bc9-bbfa-a50cc6988667
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// rg-aekxxzuad5zzzz
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// app-demo
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// Service type
	//
	// example:
	//
	// apm
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// if can be null:
	// true
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListServicesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListServicesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServicesShrinkRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServicesShrinkRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListServicesShrinkRequest) SetMaxResults(v int32) *ListServicesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesShrinkRequest) SetNextToken(v string) *ListServicesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesShrinkRequest) SetResourceGroupId(v string) *ListServicesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceName(v string) *ListServicesShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceType(v string) *ListServicesShrinkRequest {
	s.ServiceType = &v
	return s
}

func (s *ListServicesShrinkRequest) SetTagsShrink(v string) *ListServicesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListServicesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
