// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppViewTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppViewTemplateRequest
	GetAppId() *string
	SetTemplate(v *DeleteAppViewTemplateRequestTemplate) *DeleteAppViewTemplateRequest
	GetTemplate() *DeleteAppViewTemplateRequestTemplate
}

type DeleteAppViewTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Template *DeleteAppViewTemplateRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s DeleteAppViewTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppViewTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppViewTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppViewTemplateRequest) GetTemplate() *DeleteAppViewTemplateRequestTemplate {
	return s.Template
}

func (s *DeleteAppViewTemplateRequest) SetAppId(v string) *DeleteAppViewTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppViewTemplateRequest) SetTemplate(v *DeleteAppViewTemplateRequestTemplate) *DeleteAppViewTemplateRequest {
	s.Template = v
	return s
}

func (s *DeleteAppViewTemplateRequest) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAppViewTemplateRequestTemplate struct {
	// This parameter is required.
	//
	// example:
	//
	// xd4c****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteAppViewTemplateRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppViewTemplateRequestTemplate) GoString() string {
	return s.String()
}

func (s *DeleteAppViewTemplateRequestTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteAppViewTemplateRequestTemplate) SetTemplateId(v string) *DeleteAppViewTemplateRequestTemplate {
	s.TemplateId = &v
	return s
}

func (s *DeleteAppViewTemplateRequestTemplate) Validate() error {
	return dara.Validate(s)
}
