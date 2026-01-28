// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceNews interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v int64) *GrafanaWorkspaceNews
	GetDate() *int64
	SetDescription(v string) *GrafanaWorkspaceNews
	GetDescription() *string
	SetImage(v string) *GrafanaWorkspaceNews
	GetImage() *string
	SetLink(v string) *GrafanaWorkspaceNews
	GetLink() *string
	SetTitle(v string) *GrafanaWorkspaceNews
	GetTitle() *string
}

type GrafanaWorkspaceNews struct {
	Date        *int64  `json:"date,omitempty" xml:"date,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Image       *string `json:"image,omitempty" xml:"image,omitempty"`
	Link        *string `json:"link,omitempty" xml:"link,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GrafanaWorkspaceNews) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceNews) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceNews) GetDate() *int64 {
	return s.Date
}

func (s *GrafanaWorkspaceNews) GetDescription() *string {
	return s.Description
}

func (s *GrafanaWorkspaceNews) GetImage() *string {
	return s.Image
}

func (s *GrafanaWorkspaceNews) GetLink() *string {
	return s.Link
}

func (s *GrafanaWorkspaceNews) GetTitle() *string {
	return s.Title
}

func (s *GrafanaWorkspaceNews) SetDate(v int64) *GrafanaWorkspaceNews {
	s.Date = &v
	return s
}

func (s *GrafanaWorkspaceNews) SetDescription(v string) *GrafanaWorkspaceNews {
	s.Description = &v
	return s
}

func (s *GrafanaWorkspaceNews) SetImage(v string) *GrafanaWorkspaceNews {
	s.Image = &v
	return s
}

func (s *GrafanaWorkspaceNews) SetLink(v string) *GrafanaWorkspaceNews {
	s.Link = &v
	return s
}

func (s *GrafanaWorkspaceNews) SetTitle(v string) *GrafanaWorkspaceNews {
	s.Title = &v
	return s
}

func (s *GrafanaWorkspaceNews) Validate() error {
	return dara.Validate(s)
}
