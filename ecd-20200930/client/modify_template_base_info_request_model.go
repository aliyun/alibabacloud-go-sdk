// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyTemplateBaseInfoRequest
	GetDescription() *string
	SetTemplateId(v string) *ModifyTemplateBaseInfoRequest
	GetTemplateId() *string
	SetTemplateName(v string) *ModifyTemplateBaseInfoRequest
	GetTemplateName() *string
}

type ModifyTemplateBaseInfoRequest struct {
	// The template description.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b-0caoeogs88y*****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifyTemplateBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateBaseInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyTemplateBaseInfoRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyTemplateBaseInfoRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyTemplateBaseInfoRequest) SetDescription(v string) *ModifyTemplateBaseInfoRequest {
	s.Description = &v
	return s
}

func (s *ModifyTemplateBaseInfoRequest) SetTemplateId(v string) *ModifyTemplateBaseInfoRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyTemplateBaseInfoRequest) SetTemplateName(v string) *ModifyTemplateBaseInfoRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyTemplateBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
