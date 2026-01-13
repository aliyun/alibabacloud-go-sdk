// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesResult interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServicesResult
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesResult
	GetNextToken() *string
	SetServices(v *LrOrder) *ListServicesResult
	GetServices() *LrOrder
	SetTotal(v int32) *ListServicesResult
	GetTotal() *int32
}

type ListServicesResult struct {
	MaxResults *int32   `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Services   *LrOrder `json:"services,omitempty" xml:"services,omitempty"`
	Total      *int32   `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListServicesResult) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResult) GoString() string {
	return s.String()
}

func (s *ListServicesResult) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesResult) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesResult) GetServices() *LrOrder {
	return s.Services
}

func (s *ListServicesResult) GetTotal() *int32 {
	return s.Total
}

func (s *ListServicesResult) SetMaxResults(v int32) *ListServicesResult {
	s.MaxResults = &v
	return s
}

func (s *ListServicesResult) SetNextToken(v string) *ListServicesResult {
	s.NextToken = &v
	return s
}

func (s *ListServicesResult) SetServices(v *LrOrder) *ListServicesResult {
	s.Services = v
	return s
}

func (s *ListServicesResult) SetTotal(v int32) *ListServicesResult {
	s.Total = &v
	return s
}

func (s *ListServicesResult) Validate() error {
	if s.Services != nil {
		if err := s.Services.Validate(); err != nil {
			return err
		}
	}
	return nil
}
