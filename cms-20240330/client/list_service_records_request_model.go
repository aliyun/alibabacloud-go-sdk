// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceRecordsRequest
	GetNextToken() *string
	SetRecordType(v string) *ListServiceRecordsRequest
	GetRecordType() *string
	SetSearch(v string) *ListServiceRecordsRequest
	GetSearch() *string
}

type ListServiceRecordsRequest struct {
	// The maximum number of entries to return. Maximum value: 200.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The type of the linked entry. Currently supported:
	//
	// logCorrelation, which indicates application log association.
	//
	// This parameter is required.
	//
	// example:
	//
	// logCorrelation
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
	// The filter information for service-linked entries.
	//
	// example:
	//
	// {
	//
	//   "serviceName": "my-service",
	//
	//   "storeName": "my-logstore"
	//
	// }
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
}

func (s ListServiceRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceRecordsRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *ListServiceRecordsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListServiceRecordsRequest) SetMaxResults(v int32) *ListServiceRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRecordsRequest) SetNextToken(v string) *ListServiceRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceRecordsRequest) SetRecordType(v string) *ListServiceRecordsRequest {
	s.RecordType = &v
	return s
}

func (s *ListServiceRecordsRequest) SetSearch(v string) *ListServiceRecordsRequest {
	s.Search = &v
	return s
}

func (s *ListServiceRecordsRequest) Validate() error {
	return dara.Validate(s)
}
