// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppStreamingOutTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppStreamingOutTemplateShrinkRequest
	GetAppId() *string
	SetStreamingOutTemplateShrink(v string) *DeleteAppStreamingOutTemplateShrinkRequest
	GetStreamingOutTemplateShrink() *string
}

type DeleteAppStreamingOutTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplateShrink *string `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty"`
}

func (s DeleteAppStreamingOutTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppStreamingOutTemplateShrinkRequest) GetStreamingOutTemplateShrink() *string {
	return s.StreamingOutTemplateShrink
}

func (s *DeleteAppStreamingOutTemplateShrinkRequest) SetAppId(v string) *DeleteAppStreamingOutTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateShrinkRequest) SetStreamingOutTemplateShrink(v string) *DeleteAppStreamingOutTemplateShrinkRequest {
	s.StreamingOutTemplateShrink = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
