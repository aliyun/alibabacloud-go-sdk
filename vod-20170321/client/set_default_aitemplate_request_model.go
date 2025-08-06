// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultAITemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *SetDefaultAITemplateRequest
	GetTemplateId() *string
}

type SetDefaultAITemplateRequest struct {
	// The ID of the AI template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1706a0063dd733f6a823ef32e0a5****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SetDefaultAITemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultAITemplateRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultAITemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetDefaultAITemplateRequest) SetTemplateId(v string) *SetDefaultAITemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *SetDefaultAITemplateRequest) Validate() error {
	return dara.Validate(s)
}
