// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateType(v string) *ListAITemplateRequest
	GetTemplateType() *string
}

type ListAITemplateRequest struct {
	// The type of the AI template. Valid values:
	//
	// 	- **AIMediaAudit**: automated review
	//
	// 	- **AIImage**: smart thumbnail
	//
	// This parameter is required.
	//
	// example:
	//
	// AIMediaAudit
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListAITemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAITemplateRequest) GoString() string {
	return s.String()
}

func (s *ListAITemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListAITemplateRequest) SetTemplateType(v string) *ListAITemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ListAITemplateRequest) Validate() error {
	return dara.Validate(s)
}
