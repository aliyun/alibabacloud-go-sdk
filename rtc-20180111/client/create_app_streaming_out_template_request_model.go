// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppStreamingOutTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppStreamingOutTemplateRequest
	GetAppId() *string
	SetStreamingOutTemplate(v *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) *CreateAppStreamingOutTemplateRequest
	GetStreamingOutTemplate() *CreateAppStreamingOutTemplateRequestStreamingOutTemplate
}

type CreateAppStreamingOutTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplate *CreateAppStreamingOutTemplateRequestStreamingOutTemplate `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty" type:"Struct"`
}

func (s CreateAppStreamingOutTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppStreamingOutTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppStreamingOutTemplateRequest) GetStreamingOutTemplate() *CreateAppStreamingOutTemplateRequestStreamingOutTemplate {
	return s.StreamingOutTemplate
}

func (s *CreateAppStreamingOutTemplateRequest) SetAppId(v string) *CreateAppStreamingOutTemplateRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppStreamingOutTemplateRequest) SetStreamingOutTemplate(v *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) *CreateAppStreamingOutTemplateRequest {
	s.StreamingOutTemplate = v
	return s
}

func (s *CreateAppStreamingOutTemplateRequest) Validate() error {
	if s.StreamingOutTemplate != nil {
		if err := s.StreamingOutTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppStreamingOutTemplateRequestStreamingOutTemplate struct {
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

func (s CreateAppStreamingOutTemplateRequestStreamingOutTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateAppStreamingOutTemplateRequestStreamingOutTemplate) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) GetLayoutIds() []*string {
	return s.LayoutIds
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) GetName() *string {
	return s.Name
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) SetLayoutIds(v []*string) *CreateAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.LayoutIds = v
	return s
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) SetMediaEncode(v int32) *CreateAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.MediaEncode = &v
	return s
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) SetName(v string) *CreateAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.Name = &v
	return s
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) Validate() error {
	return dara.Validate(s)
}
