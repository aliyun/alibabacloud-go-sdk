// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListLogstashPluginsRequest
	GetName() *string
	SetPage(v int32) *ListLogstashPluginsRequest
	GetPage() *int32
	SetSize(v int32) *ListLogstashPluginsRequest
	GetSize() *int32
	SetSource(v string) *ListLogstashPluginsRequest
	GetSource() *string
}

type ListLogstashPluginsRequest struct {
	// USER
	//
	// example:
	//
	// logstash-filter-clone
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 10
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The returned results.
	//
	// example:
	//
	// 3
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The description of the plug-in.
	//
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListLogstashPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListLogstashPluginsRequest) GetName() *string {
	return s.Name
}

func (s *ListLogstashPluginsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListLogstashPluginsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListLogstashPluginsRequest) GetSource() *string {
	return s.Source
}

func (s *ListLogstashPluginsRequest) SetName(v string) *ListLogstashPluginsRequest {
	s.Name = &v
	return s
}

func (s *ListLogstashPluginsRequest) SetPage(v int32) *ListLogstashPluginsRequest {
	s.Page = &v
	return s
}

func (s *ListLogstashPluginsRequest) SetSize(v int32) *ListLogstashPluginsRequest {
	s.Size = &v
	return s
}

func (s *ListLogstashPluginsRequest) SetSource(v string) *ListLogstashPluginsRequest {
	s.Source = &v
	return s
}

func (s *ListLogstashPluginsRequest) Validate() error {
	return dara.Validate(s)
}
