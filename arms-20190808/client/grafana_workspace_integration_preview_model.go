// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceIntegrationPreview interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GrafanaWorkspaceIntegrationPreview
	GetId() *string
	SetImage(v string) *GrafanaWorkspaceIntegrationPreview
	GetImage() *string
	SetName(v string) *GrafanaWorkspaceIntegrationPreview
	GetName() *string
	SetThumbnail(v string) *GrafanaWorkspaceIntegrationPreview
	GetThumbnail() *string
}

type GrafanaWorkspaceIntegrationPreview struct {
	Id        *string `json:"id,omitempty" xml:"id,omitempty"`
	Image     *string `json:"image,omitempty" xml:"image,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Thumbnail *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
}

func (s GrafanaWorkspaceIntegrationPreview) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceIntegrationPreview) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceIntegrationPreview) GetId() *string {
	return s.Id
}

func (s *GrafanaWorkspaceIntegrationPreview) GetImage() *string {
	return s.Image
}

func (s *GrafanaWorkspaceIntegrationPreview) GetName() *string {
	return s.Name
}

func (s *GrafanaWorkspaceIntegrationPreview) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *GrafanaWorkspaceIntegrationPreview) SetId(v string) *GrafanaWorkspaceIntegrationPreview {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationPreview) SetImage(v string) *GrafanaWorkspaceIntegrationPreview {
	s.Image = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationPreview) SetName(v string) *GrafanaWorkspaceIntegrationPreview {
	s.Name = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationPreview) SetThumbnail(v string) *GrafanaWorkspaceIntegrationPreview {
	s.Thumbnail = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationPreview) Validate() error {
	return dara.Validate(s)
}
