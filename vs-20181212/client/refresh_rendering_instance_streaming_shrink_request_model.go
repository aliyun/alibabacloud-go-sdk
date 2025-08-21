// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshRenderingInstanceStreamingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientInfoShrink(v string) *RefreshRenderingInstanceStreamingShrinkRequest
	GetClientInfoShrink() *string
	SetRenderingInstanceId(v string) *RefreshRenderingInstanceStreamingShrinkRequest
	GetRenderingInstanceId() *string
}

type RefreshRenderingInstanceStreamingShrinkRequest struct {
	ClientInfoShrink *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s RefreshRenderingInstanceStreamingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshRenderingInstanceStreamingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshRenderingInstanceStreamingShrinkRequest) GetClientInfoShrink() *string {
	return s.ClientInfoShrink
}

func (s *RefreshRenderingInstanceStreamingShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *RefreshRenderingInstanceStreamingShrinkRequest) SetClientInfoShrink(v string) *RefreshRenderingInstanceStreamingShrinkRequest {
	s.ClientInfoShrink = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingShrinkRequest) SetRenderingInstanceId(v string) *RefreshRenderingInstanceStreamingShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *RefreshRenderingInstanceStreamingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
