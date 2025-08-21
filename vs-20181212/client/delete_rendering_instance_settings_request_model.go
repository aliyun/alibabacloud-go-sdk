// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeNames(v []*string) *DeleteRenderingInstanceSettingsRequest
	GetAttributeNames() []*string
	SetRenderingInstanceId(v string) *DeleteRenderingInstanceSettingsRequest
	GetRenderingInstanceId() *string
}

type DeleteRenderingInstanceSettingsRequest struct {
	AttributeNames []*string `json:"AttributeNames,omitempty" xml:"AttributeNames,omitempty" type:"Repeated"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DeleteRenderingInstanceSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceSettingsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceSettingsRequest) GetAttributeNames() []*string {
	return s.AttributeNames
}

func (s *DeleteRenderingInstanceSettingsRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DeleteRenderingInstanceSettingsRequest) SetAttributeNames(v []*string) *DeleteRenderingInstanceSettingsRequest {
	s.AttributeNames = v
	return s
}

func (s *DeleteRenderingInstanceSettingsRequest) SetRenderingInstanceId(v string) *DeleteRenderingInstanceSettingsRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DeleteRenderingInstanceSettingsRequest) Validate() error {
	return dara.Validate(s)
}
