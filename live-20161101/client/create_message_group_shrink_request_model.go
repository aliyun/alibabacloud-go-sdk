// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMessageGroupShrinkRequest
	GetAppId() *string
	SetCreatorId(v string) *CreateMessageGroupShrinkRequest
	GetCreatorId() *string
	SetExtensionShrink(v string) *CreateMessageGroupShrinkRequest
	GetExtensionShrink() *string
}

type CreateMessageGroupShrinkRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the creator. The ID can be up to 36 characters in length and can contain only letters and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// as****hs
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The extended field.
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateMessageGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageGroupShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMessageGroupShrinkRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateMessageGroupShrinkRequest) GetExtensionShrink() *string {
	return s.ExtensionShrink
}

func (s *CreateMessageGroupShrinkRequest) SetAppId(v string) *CreateMessageGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateMessageGroupShrinkRequest) SetCreatorId(v string) *CreateMessageGroupShrinkRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateMessageGroupShrinkRequest) SetExtensionShrink(v string) *CreateMessageGroupShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *CreateMessageGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
