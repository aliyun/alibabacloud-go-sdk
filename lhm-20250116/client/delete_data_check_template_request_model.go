// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateIds(v []*string) *DeleteDataCheckTemplateRequest
	GetTemplateIds() []*string
}

type DeleteDataCheckTemplateRequest struct {
	// The list of validation template IDs. Batch operations are supported.
	TemplateIds []*string `json:"templateIds,omitempty" xml:"templateIds,omitempty" type:"Repeated"`
}

func (s DeleteDataCheckTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckTemplateRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *DeleteDataCheckTemplateRequest) SetTemplateIds(v []*string) *DeleteDataCheckTemplateRequest {
	s.TemplateIds = v
	return s
}

func (s *DeleteDataCheckTemplateRequest) Validate() error {
	return dara.Validate(s)
}
