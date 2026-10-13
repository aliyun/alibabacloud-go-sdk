// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHiMarketConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeveloperId(v string) *ListHiMarketConsumersRequest
	GetDeveloperId() *string
	SetMaxResults(v int32) *ListHiMarketConsumersRequest
	GetMaxResults() *int32
	SetName(v string) *ListHiMarketConsumersRequest
	GetName() *string
	SetNextToken(v string) *ListHiMarketConsumersRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListHiMarketConsumersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHiMarketConsumersRequest
	GetPageSize() *int32
	SetPortalId(v string) *ListHiMarketConsumersRequest
	GetPortalId() *string
}

type ListHiMarketConsumersRequest struct {
	// The developer ID. This parameter is used to filter consumers by developer.
	//
	// example:
	//
	// dev-xxxxxxxx
	DeveloperId *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	// The maximum number of entries per page. The default value and valid values are not specified in the API definition.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The search keyword for the consumer name.
	//
	// example:
	//
	// consumer-1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token. To retrieve the next page, pass the nextToken value returned in the previous response. Do not specify this parameter for the first query.
	//
	// example:
	//
	// token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The portal ID. This parameter is used to filter consumers by portal.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
}

func (s ListHiMarketConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumersRequest) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumersRequest) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *ListHiMarketConsumersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHiMarketConsumersRequest) GetName() *string {
	return s.Name
}

func (s *ListHiMarketConsumersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHiMarketConsumersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHiMarketConsumersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHiMarketConsumersRequest) GetPortalId() *string {
	return s.PortalId
}

func (s *ListHiMarketConsumersRequest) SetDeveloperId(v string) *ListHiMarketConsumersRequest {
	s.DeveloperId = &v
	return s
}

func (s *ListHiMarketConsumersRequest) SetMaxResults(v int32) *ListHiMarketConsumersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHiMarketConsumersRequest) SetName(v string) *ListHiMarketConsumersRequest {
	s.Name = &v
	return s
}

func (s *ListHiMarketConsumersRequest) SetNextToken(v string) *ListHiMarketConsumersRequest {
	s.NextToken = &v
	return s
}

func (s *ListHiMarketConsumersRequest) SetPageNumber(v int32) *ListHiMarketConsumersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHiMarketConsumersRequest) SetPageSize(v int32) *ListHiMarketConsumersRequest {
	s.PageSize = &v
	return s
}

func (s *ListHiMarketConsumersRequest) SetPortalId(v string) *ListHiMarketConsumersRequest {
	s.PortalId = &v
	return s
}

func (s *ListHiMarketConsumersRequest) Validate() error {
	return dara.Validate(s)
}
