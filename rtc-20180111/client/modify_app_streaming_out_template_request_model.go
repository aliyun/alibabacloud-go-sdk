// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppStreamingOutTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppStreamingOutTemplateRequest
	GetAppId() *string
	SetStreamingOutTemplate(v *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) *ModifyAppStreamingOutTemplateRequest
	GetStreamingOutTemplate() *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate
}

type ModifyAppStreamingOutTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplate *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty" type:"Struct"`
}

func (s ModifyAppStreamingOutTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppStreamingOutTemplateRequest) GetStreamingOutTemplate() *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	return s.StreamingOutTemplate
}

func (s *ModifyAppStreamingOutTemplateRequest) SetAppId(v string) *ModifyAppStreamingOutTemplateRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequest) SetStreamingOutTemplate(v *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) *ModifyAppStreamingOutTemplateRequest {
	s.StreamingOutTemplate = v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequest) Validate() error {
	if s.StreamingOutTemplate != nil {
		if err := s.StreamingOutTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppStreamingOutTemplateRequestStreamingOutTemplate struct {
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

func (s ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) GetName() *string {
	return s.Name
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) SetLayoutIds(v []*string) *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.LayoutIds = v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) SetMediaEncode(v int32) *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.MediaEncode = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) SetName(v string) *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.Name = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) SetTemplateId(v string) *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.TemplateId = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) Validate() error {
	return dara.Validate(s)
}
