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
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
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
	// The list of layout IDs. A maximum of 16 layout IDs are supported.
	//
	// Note: The layouts in the list must not have the same number of data sources.
	//
	// This parameter is required.
	LayoutIds []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// The encoding options.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Template
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
