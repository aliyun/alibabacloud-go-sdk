// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeprecatedTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListDeprecatedTemplatesRequest
	GetName() *string
	SetPage(v int32) *ListDeprecatedTemplatesRequest
	GetPage() *int32
	SetSize(v int32) *ListDeprecatedTemplatesRequest
	GetSize() *int32
}

type ListDeprecatedTemplatesRequest struct {
	// example:
	//
	// component-openstore-index-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s ListDeprecatedTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeprecatedTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListDeprecatedTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListDeprecatedTemplatesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListDeprecatedTemplatesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListDeprecatedTemplatesRequest) SetName(v string) *ListDeprecatedTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListDeprecatedTemplatesRequest) SetPage(v int32) *ListDeprecatedTemplatesRequest {
	s.Page = &v
	return s
}

func (s *ListDeprecatedTemplatesRequest) SetSize(v int32) *ListDeprecatedTemplatesRequest {
	s.Size = &v
	return s
}

func (s *ListDeprecatedTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
