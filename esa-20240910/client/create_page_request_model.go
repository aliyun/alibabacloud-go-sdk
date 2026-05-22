// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreatePageRequest
	GetContent() *string
	SetContentType(v string) *CreatePageRequest
	GetContentType() *string
	SetDescription(v string) *CreatePageRequest
	GetDescription() *string
	SetName(v string) *CreatePageRequest
	GetName() *string
	SetSiteIds(v []*int64) *CreatePageRequest
	GetSiteIds() []*int64
}

type CreatePageRequest struct {
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
	// example:
	//
	// a custom deny page
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	SiteIds []*int64 `json:"SiteIds,omitempty" xml:"SiteIds,omitempty" type:"Repeated"`
}

func (s CreatePageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePageRequest) GoString() string {
	return s.String()
}

func (s *CreatePageRequest) GetContent() *string {
	return s.Content
}

func (s *CreatePageRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreatePageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePageRequest) GetName() *string {
	return s.Name
}

func (s *CreatePageRequest) GetSiteIds() []*int64 {
	return s.SiteIds
}

func (s *CreatePageRequest) SetContent(v string) *CreatePageRequest {
	s.Content = &v
	return s
}

func (s *CreatePageRequest) SetContentType(v string) *CreatePageRequest {
	s.ContentType = &v
	return s
}

func (s *CreatePageRequest) SetDescription(v string) *CreatePageRequest {
	s.Description = &v
	return s
}

func (s *CreatePageRequest) SetName(v string) *CreatePageRequest {
	s.Name = &v
	return s
}

func (s *CreatePageRequest) SetSiteIds(v []*int64) *CreatePageRequest {
	s.SiteIds = v
	return s
}

func (s *CreatePageRequest) Validate() error {
	return dara.Validate(s)
}
