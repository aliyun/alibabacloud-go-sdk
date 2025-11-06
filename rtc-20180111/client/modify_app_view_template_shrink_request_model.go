// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppViewTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppViewTemplateShrinkRequest
	GetAppId() *string
	SetTemplateShrink(v string) *ModifyAppViewTemplateShrinkRequest
	GetTemplateShrink() *string
}

type ModifyAppViewTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ModifyAppViewTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppViewTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppViewTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppViewTemplateShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *ModifyAppViewTemplateShrinkRequest) SetAppId(v string) *ModifyAppViewTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppViewTemplateShrinkRequest) SetTemplateShrink(v string) *ModifyAppViewTemplateShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *ModifyAppViewTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
