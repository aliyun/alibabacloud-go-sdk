// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMigrationZonesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMigrationZonesRequest
	GetNextToken() *string
}

type ListMigrationZonesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMigrationZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationZonesRequest) GoString() string {
	return s.String()
}

func (s *ListMigrationZonesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMigrationZonesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMigrationZonesRequest) SetMaxResults(v int32) *ListMigrationZonesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMigrationZonesRequest) SetNextToken(v string) *ListMigrationZonesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMigrationZonesRequest) Validate() error {
	return dara.Validate(s)
}
