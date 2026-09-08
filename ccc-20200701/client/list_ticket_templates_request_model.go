// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ListTicketTemplatesRequest
	GetCategoryId() *string
	SetInstanceId(v string) *ListTicketTemplatesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListTicketTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTicketTemplatesRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListTicketTemplatesRequest
	GetSearchPattern() *string
	SetState(v string) *ListTicketTemplatesRequest
	GetState() *string
}

type ListTicketTemplatesRequest struct {
	// The ID of the ticket category.
	//
	// example:
	//
	// 43c2671b-********86d0-6bd187905cc8
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value must be between 1 and 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value must be between 1 and 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the template. Fuzzy search is supported.
	//
	// example:
	//
	// 模板1
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	// The status of the template.
	//
	// - Enabled: The template is published.
	//
	// - Disabled: The template is unpublished.
	//
	// example:
	//
	// Enabled
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListTicketTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTicketTemplatesRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListTicketTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTicketTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTicketTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketTemplatesRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListTicketTemplatesRequest) GetState() *string {
	return s.State
}

func (s *ListTicketTemplatesRequest) SetCategoryId(v string) *ListTicketTemplatesRequest {
	s.CategoryId = &v
	return s
}

func (s *ListTicketTemplatesRequest) SetInstanceId(v string) *ListTicketTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTicketTemplatesRequest) SetPageNumber(v int32) *ListTicketTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTicketTemplatesRequest) SetPageSize(v int32) *ListTicketTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketTemplatesRequest) SetSearchPattern(v string) *ListTicketTemplatesRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListTicketTemplatesRequest) SetState(v string) *ListTicketTemplatesRequest {
	s.State = &v
	return s
}

func (s *ListTicketTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
