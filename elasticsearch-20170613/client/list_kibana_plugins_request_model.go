// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKibanaPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v string) *ListKibanaPluginsRequest
	GetPage() *string
	SetSize(v int32) *ListKibanaPluginsRequest
	GetSize() *int32
}

type ListKibanaPluginsRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	Page *string `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListKibanaPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsRequest) GetPage() *string {
	return s.Page
}

func (s *ListKibanaPluginsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListKibanaPluginsRequest) SetPage(v string) *ListKibanaPluginsRequest {
	s.Page = &v
	return s
}

func (s *ListKibanaPluginsRequest) SetSize(v int32) *ListKibanaPluginsRequest {
	s.Size = &v
	return s
}

func (s *ListKibanaPluginsRequest) Validate() error {
	return dara.Validate(s)
}
