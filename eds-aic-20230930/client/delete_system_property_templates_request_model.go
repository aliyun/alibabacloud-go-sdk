// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSystemPropertyTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateIds(v []*string) *DeleteSystemPropertyTemplatesRequest
	GetTemplateIds() []*string
}

type DeleteSystemPropertyTemplatesRequest struct {
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s DeleteSystemPropertyTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSystemPropertyTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteSystemPropertyTemplatesRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *DeleteSystemPropertyTemplatesRequest) SetTemplateIds(v []*string) *DeleteSystemPropertyTemplatesRequest {
	s.TemplateIds = v
	return s
}

func (s *DeleteSystemPropertyTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
