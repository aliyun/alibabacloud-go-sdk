// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListPluginsRequest
	GetName() *string
	SetPage(v string) *ListPluginsRequest
	GetPage() *string
	SetSize(v int32) *ListPluginsRequest
	GetSize() *int32
	SetSource(v string) *ListPluginsRequest
	GetSource() *string
}

type ListPluginsRequest struct {
	// SYSTEM
	//
	// example:
	//
	// analysis-ik
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1
	Page *string `json:"page,omitempty" xml:"page,omitempty"`
	// The header of the response.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// SYSTEM
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListPluginsRequest) GetName() *string {
	return s.Name
}

func (s *ListPluginsRequest) GetPage() *string {
	return s.Page
}

func (s *ListPluginsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListPluginsRequest) GetSource() *string {
	return s.Source
}

func (s *ListPluginsRequest) SetName(v string) *ListPluginsRequest {
	s.Name = &v
	return s
}

func (s *ListPluginsRequest) SetPage(v string) *ListPluginsRequest {
	s.Page = &v
	return s
}

func (s *ListPluginsRequest) SetSize(v int32) *ListPluginsRequest {
	s.Size = &v
	return s
}

func (s *ListPluginsRequest) SetSource(v string) *ListPluginsRequest {
	s.Source = &v
	return s
}

func (s *ListPluginsRequest) Validate() error {
	return dara.Validate(s)
}
