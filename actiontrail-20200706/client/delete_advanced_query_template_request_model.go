// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdvancedQueryTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteAdvancedQueryTemplateRequest
	GetTemplateId() *string
}

type DeleteAdvancedQueryTemplateRequest struct {
	// example:
	//
	// utpl-QNL3dpYkQcyjZxrIQCciqQ
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteAdvancedQueryTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdvancedQueryTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAdvancedQueryTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteAdvancedQueryTemplateRequest) SetTemplateId(v string) *DeleteAdvancedQueryTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteAdvancedQueryTemplateRequest) Validate() error {
	return dara.Validate(s)
}
