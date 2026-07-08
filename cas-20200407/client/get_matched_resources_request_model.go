// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMatchedResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIds(v string) *GetMatchedResourcesRequest
	GetCertIds() *string
	SetMaxResults(v int32) *GetMatchedResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetMatchedResourcesRequest
	GetNextToken() *string
	SetResourceScope(v string) *GetMatchedResourcesRequest
	GetResourceScope() *string
}

type GetMatchedResourcesRequest struct {
	// The certificate IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19736665
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// Because of the large number of matched resources, the backend uses aggregation and does not support pagination. This parameter is reserved. By default, a maximum of 2,000 entries are returned.
	//
	// example:
	//
	// 2000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Because of the large number of matched resources, the backend uses aggregation and does not support pagination. This parameter is reserved.
	//
	// example:
	//
	// 666
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource match scope. This parameter can be empty.
	//
	// example:
	//
	// 目前支持全部云产品
	ResourceScope *string `json:"ResourceScope,omitempty" xml:"ResourceScope,omitempty"`
}

func (s GetMatchedResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMatchedResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetMatchedResourcesRequest) GetCertIds() *string {
	return s.CertIds
}

func (s *GetMatchedResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetMatchedResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetMatchedResourcesRequest) GetResourceScope() *string {
	return s.ResourceScope
}

func (s *GetMatchedResourcesRequest) SetCertIds(v string) *GetMatchedResourcesRequest {
	s.CertIds = &v
	return s
}

func (s *GetMatchedResourcesRequest) SetMaxResults(v int32) *GetMatchedResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *GetMatchedResourcesRequest) SetNextToken(v string) *GetMatchedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *GetMatchedResourcesRequest) SetResourceScope(v string) *GetMatchedResourcesRequest {
	s.ResourceScope = &v
	return s
}

func (s *GetMatchedResourcesRequest) Validate() error {
	return dara.Validate(s)
}
