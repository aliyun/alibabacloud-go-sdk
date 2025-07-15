// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDocumentsRequest
	GetInstanceId() *string
	SetNextPageToken(v string) *ListDocumentsRequest
	GetNextPageToken() *string
	SetPageSize(v int32) *ListDocumentsRequest
	GetPageSize() *int32
	SetRequestId(v string) *ListDocumentsRequest
	GetRequestId() *string
	SetSchemaId(v string) *ListDocumentsRequest
	GetSchemaId() *string
	SetSearchPattern(v string) *ListDocumentsRequest
	GetSearchPattern() *string
	SetSorts(v []*ListDocumentsRequestSorts) *ListDocumentsRequest
	GetSorts() []*ListDocumentsRequestSorts
}

type ListDocumentsRequest struct {
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
	SchemaId      *string                      `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	SearchPattern *string                      `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	Sorts         []*ListDocumentsRequestSorts `json:"Sorts,omitempty" xml:"Sorts,omitempty" type:"Repeated"`
}

func (s ListDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDocumentsRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDocumentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDocumentsRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentsRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListDocumentsRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListDocumentsRequest) GetSorts() []*ListDocumentsRequestSorts {
	return s.Sorts
}

func (s *ListDocumentsRequest) SetInstanceId(v string) *ListDocumentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDocumentsRequest) SetNextPageToken(v string) *ListDocumentsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListDocumentsRequest) SetPageSize(v int32) *ListDocumentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDocumentsRequest) SetRequestId(v string) *ListDocumentsRequest {
	s.RequestId = &v
	return s
}

func (s *ListDocumentsRequest) SetSchemaId(v string) *ListDocumentsRequest {
	s.SchemaId = &v
	return s
}

func (s *ListDocumentsRequest) SetSearchPattern(v string) *ListDocumentsRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListDocumentsRequest) SetSorts(v []*ListDocumentsRequestSorts) *ListDocumentsRequest {
	s.Sorts = v
	return s
}

func (s *ListDocumentsRequest) Validate() error {
	return dara.Validate(s)
}

type ListDocumentsRequestSorts struct {
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// name
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
}

func (s ListDocumentsRequestSorts) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsRequestSorts) GoString() string {
	return s.String()
}

func (s *ListDocumentsRequestSorts) GetOrder() *string {
	return s.Order
}

func (s *ListDocumentsRequestSorts) GetPropertyName() *string {
	return s.PropertyName
}

func (s *ListDocumentsRequestSorts) SetOrder(v string) *ListDocumentsRequestSorts {
	s.Order = &v
	return s
}

func (s *ListDocumentsRequestSorts) SetPropertyName(v string) *ListDocumentsRequestSorts {
	s.PropertyName = &v
	return s
}

func (s *ListDocumentsRequestSorts) Validate() error {
	return dara.Validate(s)
}
