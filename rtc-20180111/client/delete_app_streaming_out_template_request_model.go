// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppStreamingOutTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppStreamingOutTemplateRequest
	GetAppId() *string
	SetStreamingOutTemplate(v *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) *DeleteAppStreamingOutTemplateRequest
	GetStreamingOutTemplate() *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate
}

type DeleteAppStreamingOutTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplate *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty" type:"Struct"`
}

func (s DeleteAppStreamingOutTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppStreamingOutTemplateRequest) GetStreamingOutTemplate() *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate {
	return s.StreamingOutTemplate
}

func (s *DeleteAppStreamingOutTemplateRequest) SetAppId(v string) *DeleteAppStreamingOutTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateRequest) SetStreamingOutTemplate(v *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) *DeleteAppStreamingOutTemplateRequest {
	s.StreamingOutTemplate = v
	return s
}

func (s *DeleteAppStreamingOutTemplateRequest) Validate() error {
	if s.StreamingOutTemplate != nil {
		if err := s.StreamingOutTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAppStreamingOutTemplateRequestStreamingOutTemplate struct {
	// This parameter is required.
	//
	// example:
	//
	// xd4c****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) SetTemplateId(v string) *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.TemplateId = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) Validate() error {
	return dara.Validate(s)
}
