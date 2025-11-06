// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppViewTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppViewTemplateShrinkRequest
	GetAppId() *string
	SetTemplateShrink(v string) *DeleteAppViewTemplateShrinkRequest
	GetTemplateShrink() *string
}

type DeleteAppViewTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DeleteAppViewTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppViewTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppViewTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppViewTemplateShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *DeleteAppViewTemplateShrinkRequest) SetAppId(v string) *DeleteAppViewTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppViewTemplateShrinkRequest) SetTemplateShrink(v string) *DeleteAppViewTemplateShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *DeleteAppViewTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
