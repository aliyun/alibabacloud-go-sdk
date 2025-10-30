// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesRequest
	GetNextToken() *string
	SetServiceType(v string) *ListServicesRequest
	GetServiceType() *string
}

type ListServicesRequest struct {
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
	// Service type
	//
	// example:
	//
	// apm
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesRequest) SetMaxResults(v int32) *ListServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetServiceType(v string) *ListServicesRequest {
	s.ServiceType = &v
	return s
}

func (s *ListServicesRequest) Validate() error {
	return dara.Validate(s)
}
