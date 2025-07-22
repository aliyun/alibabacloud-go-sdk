// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppStreamingOutTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppStreamingOutTemplateShrinkRequest
	GetAppId() *string
	SetStreamingOutTemplateShrink(v string) *ModifyAppStreamingOutTemplateShrinkRequest
	GetStreamingOutTemplateShrink() *string
}

type ModifyAppStreamingOutTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplateShrink *string `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty"`
}

func (s ModifyAppStreamingOutTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppStreamingOutTemplateShrinkRequest) GetStreamingOutTemplateShrink() *string {
	return s.StreamingOutTemplateShrink
}

func (s *ModifyAppStreamingOutTemplateShrinkRequest) SetAppId(v string) *ModifyAppStreamingOutTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateShrinkRequest) SetStreamingOutTemplateShrink(v string) *ModifyAppStreamingOutTemplateShrinkRequest {
	s.StreamingOutTemplateShrink = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
