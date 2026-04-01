// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListUserPluginRequest
	GetName() *string
	SetPage(v string) *ListUserPluginRequest
	GetPage() *string
	SetSize(v string) *ListUserPluginRequest
	GetSize() *string
}

type ListUserPluginRequest struct {
	// example:
	//
	// my-plugin
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Page *string `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 50
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListUserPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserPluginRequest) GoString() string {
	return s.String()
}

func (s *ListUserPluginRequest) GetName() *string {
	return s.Name
}

func (s *ListUserPluginRequest) GetPage() *string {
	return s.Page
}

func (s *ListUserPluginRequest) GetSize() *string {
	return s.Size
}

func (s *ListUserPluginRequest) SetName(v string) *ListUserPluginRequest {
	s.Name = &v
	return s
}

func (s *ListUserPluginRequest) SetPage(v string) *ListUserPluginRequest {
	s.Page = &v
	return s
}

func (s *ListUserPluginRequest) SetSize(v string) *ListUserPluginRequest {
	s.Size = &v
	return s
}

func (s *ListUserPluginRequest) Validate() error {
	return dara.Validate(s)
}
