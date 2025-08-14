// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateIds(v string) *DeleteTemplateRequest
	GetTemplateIds() *string
}

type DeleteTemplateRequest struct {
	// The IDs of the templates that you want to delete. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****,****20b48fb04483915d4f2cd8ac****
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) GetTemplateIds() *string {
	return s.TemplateIds
}

func (s *DeleteTemplateRequest) SetTemplateIds(v string) *DeleteTemplateRequest {
	s.TemplateIds = &v
	return s
}

func (s *DeleteTemplateRequest) Validate() error {
	return dara.Validate(s)
}
