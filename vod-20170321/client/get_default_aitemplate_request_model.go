// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultAITemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateType(v string) *GetDefaultAITemplateRequest
	GetTemplateType() *string
}

type GetDefaultAITemplateRequest struct {
	// The type of the AI template. Set the value to **AIMediaAudit**, which specifies the automated review.
	//
	// This parameter is required.
	//
	// example:
	//
	// AIMediaAudit
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetDefaultAITemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultAITemplateRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultAITemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetDefaultAITemplateRequest) SetTemplateType(v string) *GetDefaultAITemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *GetDefaultAITemplateRequest) Validate() error {
	return dara.Validate(s)
}
