// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDocumentsShrinkRequest
	GetInstanceId() *string
	SetNextPageToken(v string) *ListDocumentsShrinkRequest
	GetNextPageToken() *string
	SetPageSize(v int32) *ListDocumentsShrinkRequest
	GetPageSize() *int32
	SetRequestId(v string) *ListDocumentsShrinkRequest
	GetRequestId() *string
	SetSchemaId(v string) *ListDocumentsShrinkRequest
	GetSchemaId() *string
	SetSearchPattern(v string) *ListDocumentsShrinkRequest
	GetSearchPattern() *string
	SetSortsShrink(v string) *ListDocumentsShrinkRequest
	GetSortsShrink() *string
}

type ListDocumentsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d278629c-c687-4aa3-b044-4fe9b012e7ef
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// eyJ0YWJsZUlkIjoiY2Y2MTQxYjA5NDY0NDUxMzk5YjFjMTA5YTMxZWNkMzEiLCJ0b2tlbiI6IjAwMDAwMDAwMDAwNzAzNzcifQ==
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// EAF3C248-E123-441B-A545-B6CD02E98EED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// schema id
	//
	// This parameter is required.
	//
	// example:
	//
	// profile
	SchemaId      *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	SortsShrink   *string `json:"Sorts,omitempty" xml:"Sorts,omitempty"`
}

func (s ListDocumentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDocumentsShrinkRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDocumentsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDocumentsShrinkRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentsShrinkRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListDocumentsShrinkRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListDocumentsShrinkRequest) GetSortsShrink() *string {
	return s.SortsShrink
}

func (s *ListDocumentsShrinkRequest) SetInstanceId(v string) *ListDocumentsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDocumentsShrinkRequest) SetNextPageToken(v string) *ListDocumentsShrinkRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListDocumentsShrinkRequest) SetPageSize(v int32) *ListDocumentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDocumentsShrinkRequest) SetRequestId(v string) *ListDocumentsShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *ListDocumentsShrinkRequest) SetSchemaId(v string) *ListDocumentsShrinkRequest {
	s.SchemaId = &v
	return s
}

func (s *ListDocumentsShrinkRequest) SetSearchPattern(v string) *ListDocumentsShrinkRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListDocumentsShrinkRequest) SetSortsShrink(v string) *ListDocumentsShrinkRequest {
	s.SortsShrink = &v
	return s
}

func (s *ListDocumentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
