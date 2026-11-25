// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddManifestEditTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddManifestEditTemplateRequest
	GetName() *string
	SetTemplates(v string) *AddManifestEditTemplateRequest
	GetTemplates() *string
}

type AddManifestEditTemplateRequest struct {
	// The name of the template. Spaces are allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// add-multilang-tracks
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The video encoding configuration, such as `{"Codec":"H.264","Bitrate":"5000000"}`.
	//
	// example:
	//
	// {"Codec":"H.264","Bitrate":"5000000"}
	Templates *string `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s AddManifestEditTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddManifestEditTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddManifestEditTemplateRequest) GetName() *string {
	return s.Name
}

func (s *AddManifestEditTemplateRequest) GetTemplates() *string {
	return s.Templates
}

func (s *AddManifestEditTemplateRequest) SetName(v string) *AddManifestEditTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddManifestEditTemplateRequest) SetTemplates(v string) *AddManifestEditTemplateRequest {
	s.Templates = &v
	return s
}

func (s *AddManifestEditTemplateRequest) Validate() error {
	return dara.Validate(s)
}
