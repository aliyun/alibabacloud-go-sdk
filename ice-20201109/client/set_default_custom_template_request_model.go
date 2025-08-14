// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *SetDefaultCustomTemplateRequest
	GetTemplateId() *string
}

type SetDefaultCustomTemplateRequest struct {
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SetDefaultCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultCustomTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetDefaultCustomTemplateRequest) SetTemplateId(v string) *SetDefaultCustomTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *SetDefaultCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
