// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreatePageShrinkRequest
	GetContent() *string
	SetContentType(v string) *CreatePageShrinkRequest
	GetContentType() *string
	SetDescription(v string) *CreatePageShrinkRequest
	GetDescription() *string
	SetName(v string) *CreatePageShrinkRequest
	GetName() *string
	SetSiteIdsShrink(v string) *CreatePageShrinkRequest
	GetSiteIdsShrink() *string
}

type CreatePageShrinkRequest struct {
	// The Base64-encoded page content. Example: "PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=", which indicates "hello page".
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The Content-Type field in the HTTP header. Valid values:
	//
	// 	- text/html
	//
	// 	- application/json
	//
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The description of the page.
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom error page.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SiteIdsShrink *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
}

func (s CreatePageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePageShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreatePageShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreatePageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePageShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreatePageShrinkRequest) GetSiteIdsShrink() *string {
	return s.SiteIdsShrink
}

func (s *CreatePageShrinkRequest) SetContent(v string) *CreatePageShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreatePageShrinkRequest) SetContentType(v string) *CreatePageShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *CreatePageShrinkRequest) SetDescription(v string) *CreatePageShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePageShrinkRequest) SetName(v string) *CreatePageShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreatePageShrinkRequest) SetSiteIdsShrink(v string) *CreatePageShrinkRequest {
	s.SiteIdsShrink = &v
	return s
}

func (s *CreatePageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
