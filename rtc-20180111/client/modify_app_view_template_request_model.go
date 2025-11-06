// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppViewTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppViewTemplateRequest
	GetAppId() *string
	SetTemplate(v *ModifyAppViewTemplateRequestTemplate) *ModifyAppViewTemplateRequest
	GetTemplate() *ModifyAppViewTemplateRequestTemplate
}

type ModifyAppViewTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Template *ModifyAppViewTemplateRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s ModifyAppViewTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppViewTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppViewTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppViewTemplateRequest) GetTemplate() *ModifyAppViewTemplateRequestTemplate {
	return s.Template
}

func (s *ModifyAppViewTemplateRequest) SetAppId(v string) *ModifyAppViewTemplateRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppViewTemplateRequest) SetTemplate(v *ModifyAppViewTemplateRequestTemplate) *ModifyAppViewTemplateRequest {
	s.Template = v
	return s
}

func (s *ModifyAppViewTemplateRequest) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppViewTemplateRequestTemplate struct {
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
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyAppViewTemplateRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppViewTemplateRequestTemplate) GoString() string {
	return s.String()
}

func (s *ModifyAppViewTemplateRequestTemplate) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *ModifyAppViewTemplateRequestTemplate) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *ModifyAppViewTemplateRequestTemplate) GetName() *string {
	return s.Name
}

func (s *ModifyAppViewTemplateRequestTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyAppViewTemplateRequestTemplate) SetLayoutIds(v []*string) *ModifyAppViewTemplateRequestTemplate {
	s.LayoutIds = v
	return s
}

func (s *ModifyAppViewTemplateRequestTemplate) SetMediaEncode(v int32) *ModifyAppViewTemplateRequestTemplate {
	s.MediaEncode = &v
	return s
}

func (s *ModifyAppViewTemplateRequestTemplate) SetName(v string) *ModifyAppViewTemplateRequestTemplate {
	s.Name = &v
	return s
}

func (s *ModifyAppViewTemplateRequestTemplate) SetTemplateId(v string) *ModifyAppViewTemplateRequestTemplate {
	s.TemplateId = &v
	return s
}

func (s *ModifyAppViewTemplateRequestTemplate) Validate() error {
	return dara.Validate(s)
}
