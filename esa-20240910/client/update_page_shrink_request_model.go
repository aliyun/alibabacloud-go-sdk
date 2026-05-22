// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdatePageShrinkRequest
	GetContent() *string
	SetContentType(v string) *UpdatePageShrinkRequest
	GetContentType() *string
	SetDescription(v string) *UpdatePageShrinkRequest
	GetDescription() *string
	SetId(v int64) *UpdatePageShrinkRequest
	GetId() *int64
	SetName(v string) *UpdatePageShrinkRequest
	GetName() *string
	SetSiteIdsShrink(v string) *UpdatePageShrinkRequest
	GetSiteIdsShrink() *string
}

type UpdatePageShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PGh0bWw+aGVsbG8gcGFnZTwvaHRtbD4=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text/html
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SiteIdsShrink *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
}

func (s UpdatePageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePageShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *UpdatePageShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UpdatePageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePageShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdatePageShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePageShrinkRequest) GetSiteIdsShrink() *string {
	return s.SiteIdsShrink
}

func (s *UpdatePageShrinkRequest) SetContent(v string) *UpdatePageShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetContentType(v string) *UpdatePageShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetDescription(v string) *UpdatePageShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetId(v int64) *UpdatePageShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetName(v string) *UpdatePageShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdatePageShrinkRequest) SetSiteIdsShrink(v string) *UpdatePageShrinkRequest {
	s.SiteIdsShrink = &v
	return s
}

func (s *UpdatePageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
