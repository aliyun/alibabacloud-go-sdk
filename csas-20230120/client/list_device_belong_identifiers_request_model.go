// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBelongIdentifiersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDeviceBelongIdentifiersRequest
	GetCurrentPage() *int32
	SetIdentifierIds(v []*string) *ListDeviceBelongIdentifiersRequest
	GetIdentifierIds() []*string
	SetIdentifierType(v string) *ListDeviceBelongIdentifiersRequest
	GetIdentifierType() *string
	SetMaxResults(v int32) *ListDeviceBelongIdentifiersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeviceBelongIdentifiersRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListDeviceBelongIdentifiersRequest
	GetPageSize() *int32
}

type ListDeviceBelongIdentifiersRequest struct {
	// This parameter is required.
	CurrentPage    *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	IdentifierIds  []*string `json:"IdentifierIds,omitempty" xml:"IdentifierIds,omitempty" type:"Repeated"`
	IdentifierType *string   `json:"IdentifierType,omitempty" xml:"IdentifierType,omitempty"`
	MaxResults     *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceBelongIdentifiersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBelongIdentifiersRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceBelongIdentifiersRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDeviceBelongIdentifiersRequest) GetIdentifierIds() []*string {
	return s.IdentifierIds
}

func (s *ListDeviceBelongIdentifiersRequest) GetIdentifierType() *string {
	return s.IdentifierType
}

func (s *ListDeviceBelongIdentifiersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeviceBelongIdentifiersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeviceBelongIdentifiersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeviceBelongIdentifiersRequest) SetCurrentPage(v int32) *ListDeviceBelongIdentifiersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetIdentifierIds(v []*string) *ListDeviceBelongIdentifiersRequest {
	s.IdentifierIds = v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetIdentifierType(v string) *ListDeviceBelongIdentifiersRequest {
	s.IdentifierType = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetMaxResults(v int32) *ListDeviceBelongIdentifiersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetNextToken(v string) *ListDeviceBelongIdentifiersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetPageSize(v int32) *ListDeviceBelongIdentifiersRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) Validate() error {
	return dara.Validate(s)
}
