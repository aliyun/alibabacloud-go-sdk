// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataEventServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDataEventServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataEventServicesRequest
	GetNextToken() *string
}

type ListDataEventServicesRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// VjE6dLbnNpVmbz06****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDataEventServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventServicesRequest) GoString() string {
	return s.String()
}

func (s *ListDataEventServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataEventServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataEventServicesRequest) SetMaxResults(v int32) *ListDataEventServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataEventServicesRequest) SetNextToken(v string) *ListDataEventServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataEventServicesRequest) Validate() error {
	return dara.Validate(s)
}
