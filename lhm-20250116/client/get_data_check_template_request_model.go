// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetDataCheckTemplateRequest
	GetTemplateId() *string
}

type GetDataCheckTemplateRequest struct {
	// The check template ID (logical foreign key) that uniquely identifies a check template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s GetDataCheckTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetDataCheckTemplateRequest) SetTemplateId(v string) *GetDataCheckTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetDataCheckTemplateRequest) Validate() error {
	return dara.Validate(s)
}
