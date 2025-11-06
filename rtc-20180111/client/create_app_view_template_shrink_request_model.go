// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppViewTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppViewTemplateShrinkRequest
	GetAppId() *string
	SetTemplateShrink(v string) *CreateAppViewTemplateShrinkRequest
	GetTemplateShrink() *string
}

type CreateAppViewTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s CreateAppViewTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppViewTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppViewTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppViewTemplateShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *CreateAppViewTemplateShrinkRequest) SetAppId(v string) *CreateAppViewTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppViewTemplateShrinkRequest) SetTemplateShrink(v string) *CreateAppViewTemplateShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *CreateAppViewTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
