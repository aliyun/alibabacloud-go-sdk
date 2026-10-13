// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevelopersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDevelopersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDevelopersRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListDevelopersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDevelopersRequest
	GetPageSize() *int32
	SetPortalId(v string) *ListDevelopersRequest
	GetPortalId() *string
	SetStatus(v string) *ListDevelopersRequest
	GetStatus() *string
	SetUsername(v string) *ListDevelopersRequest
	GetUsername() *string
}

type ListDevelopersRequest struct {
	// The maximum number of entries per page. The default value and valid values are not specified in the API definition.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination cursor. Pass the nextToken value returned in the previous response to retrieve the next page. Do not pass this parameter for the first query.
	//
	// example:
	//
	// token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the portal used to filter developers.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
	// The status of the developer used to filter developers. Valid values: PENDING, APPROVED, and REJECTED.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The logon username of the developer used to filter developers.
	//
	// example:
	//
	// developer-demo
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListDevelopersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDevelopersRequest) GoString() string {
	return s.String()
}

func (s *ListDevelopersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDevelopersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDevelopersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDevelopersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDevelopersRequest) GetPortalId() *string {
	return s.PortalId
}

func (s *ListDevelopersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDevelopersRequest) GetUsername() *string {
	return s.Username
}

func (s *ListDevelopersRequest) SetMaxResults(v int32) *ListDevelopersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDevelopersRequest) SetNextToken(v string) *ListDevelopersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDevelopersRequest) SetPageNumber(v int32) *ListDevelopersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevelopersRequest) SetPageSize(v int32) *ListDevelopersRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevelopersRequest) SetPortalId(v string) *ListDevelopersRequest {
	s.PortalId = &v
	return s
}

func (s *ListDevelopersRequest) SetStatus(v string) *ListDevelopersRequest {
	s.Status = &v
	return s
}

func (s *ListDevelopersRequest) SetUsername(v string) *ListDevelopersRequest {
	s.Username = &v
	return s
}

func (s *ListDevelopersRequest) Validate() error {
	return dara.Validate(s)
}
