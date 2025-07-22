// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppStreamingOutTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppStreamingOutTemplateShrinkRequest
	GetAppId() *string
	SetStreamingOutTemplateShrink(v string) *CreateAppStreamingOutTemplateShrinkRequest
	GetStreamingOutTemplateShrink() *string
}

type CreateAppStreamingOutTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplateShrink *string `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty"`
}

func (s CreateAppStreamingOutTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppStreamingOutTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppStreamingOutTemplateShrinkRequest) GetStreamingOutTemplateShrink() *string {
	return s.StreamingOutTemplateShrink
}

func (s *CreateAppStreamingOutTemplateShrinkRequest) SetAppId(v string) *CreateAppStreamingOutTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppStreamingOutTemplateShrinkRequest) SetStreamingOutTemplateShrink(v string) *CreateAppStreamingOutTemplateShrinkRequest {
	s.StreamingOutTemplateShrink = &v
	return s
}

func (s *CreateAppStreamingOutTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
