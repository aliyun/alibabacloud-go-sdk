// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppViewTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppViewTemplateRequest
	GetAppId() *string
	SetTemplate(v *CreateAppViewTemplateRequestTemplate) *CreateAppViewTemplateRequest
	GetTemplate() *CreateAppViewTemplateRequestTemplate
}

type CreateAppViewTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Template *CreateAppViewTemplateRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s CreateAppViewTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppViewTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAppViewTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppViewTemplateRequest) GetTemplate() *CreateAppViewTemplateRequestTemplate {
	return s.Template
}

func (s *CreateAppViewTemplateRequest) SetAppId(v string) *CreateAppViewTemplateRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppViewTemplateRequest) SetTemplate(v *CreateAppViewTemplateRequestTemplate) *CreateAppViewTemplateRequest {
	s.Template = v
	return s
}

func (s *CreateAppViewTemplateRequest) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppViewTemplateRequestTemplate struct {
	// This parameter is required.
	LayoutIds []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 模版
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateAppViewTemplateRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateAppViewTemplateRequestTemplate) GoString() string {
	return s.String()
}

func (s *CreateAppViewTemplateRequestTemplate) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *CreateAppViewTemplateRequestTemplate) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *CreateAppViewTemplateRequestTemplate) GetName() *string {
	return s.Name
}

func (s *CreateAppViewTemplateRequestTemplate) SetLayoutIds(v []*string) *CreateAppViewTemplateRequestTemplate {
	s.LayoutIds = v
	return s
}

func (s *CreateAppViewTemplateRequestTemplate) SetMediaEncode(v int32) *CreateAppViewTemplateRequestTemplate {
	s.MediaEncode = &v
	return s
}

func (s *CreateAppViewTemplateRequestTemplate) SetName(v string) *CreateAppViewTemplateRequestTemplate {
	s.Name = &v
	return s
}

func (s *CreateAppViewTemplateRequestTemplate) Validate() error {
	return dara.Validate(s)
}
