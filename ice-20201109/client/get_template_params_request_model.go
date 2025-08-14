// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetTemplateParamsRequest
	GetTemplateId() *string
}

type GetTemplateParamsRequest struct {
	// The template ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParamsRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateParamsRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateParamsRequest) SetTemplateId(v string) *GetTemplateParamsRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateParamsRequest) Validate() error {
	return dara.Validate(s)
}
