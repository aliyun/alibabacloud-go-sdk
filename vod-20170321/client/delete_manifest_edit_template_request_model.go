// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteManifestEditTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteManifestEditTemplateRequest
	GetTemplateId() *string
}

type DeleteManifestEditTemplateRequest struct {
	// The ID of the ManifestEdit template to be deleted. Only letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9c6581d59240de5e258ffa4c71a3e01b
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteManifestEditTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteManifestEditTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteManifestEditTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteManifestEditTemplateRequest) SetTemplateId(v string) *DeleteManifestEditTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteManifestEditTemplateRequest) Validate() error {
	return dara.Validate(s)
}
