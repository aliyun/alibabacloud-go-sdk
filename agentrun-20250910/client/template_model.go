// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateID(v string) *Template
	GetTemplateID() *string
	SetTemplateVersion(v string) *Template
	GetTemplateVersion() *string
}

type Template struct {
	TemplateID      *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s Template) String() string {
	return dara.Prettify(s)
}

func (s Template) GoString() string {
	return s.String()
}

func (s *Template) GetTemplateID() *string {
	return s.TemplateID
}

func (s *Template) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *Template) SetTemplateID(v string) *Template {
	s.TemplateID = &v
	return s
}

func (s *Template) SetTemplateVersion(v string) *Template {
	s.TemplateVersion = &v
	return s
}

func (s *Template) Validate() error {
	return dara.Validate(s)
}
