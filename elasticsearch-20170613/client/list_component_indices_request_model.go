// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentIndicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListComponentIndicesRequest
	GetName() *string
	SetPage(v int32) *ListComponentIndicesRequest
	GetPage() *int32
	SetSize(v int32) *ListComponentIndicesRequest
	GetSize() *int32
}

type ListComponentIndicesRequest struct {
	// example:
	//
	// template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 5
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListComponentIndicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComponentIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListComponentIndicesRequest) GetName() *string {
	return s.Name
}

func (s *ListComponentIndicesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListComponentIndicesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListComponentIndicesRequest) SetName(v string) *ListComponentIndicesRequest {
	s.Name = &v
	return s
}

func (s *ListComponentIndicesRequest) SetPage(v int32) *ListComponentIndicesRequest {
	s.Page = &v
	return s
}

func (s *ListComponentIndicesRequest) SetSize(v int32) *ListComponentIndicesRequest {
	s.Size = &v
	return s
}

func (s *ListComponentIndicesRequest) Validate() error {
	return dara.Validate(s)
}
