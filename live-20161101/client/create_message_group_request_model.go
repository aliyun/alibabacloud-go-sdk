// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMessageGroupRequest
	GetAppId() *string
	SetCreatorId(v string) *CreateMessageGroupRequest
	GetCreatorId() *string
	SetExtension(v map[string]*string) *CreateMessageGroupRequest
	GetExtension() map[string]*string
}

type CreateMessageGroupRequest struct {
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
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMessageGroupRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateMessageGroupRequest) GetExtension() map[string]*string {
	return s.Extension
}

func (s *CreateMessageGroupRequest) SetAppId(v string) *CreateMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *CreateMessageGroupRequest) SetCreatorId(v string) *CreateMessageGroupRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateMessageGroupRequest) SetExtension(v map[string]*string) *CreateMessageGroupRequest {
	s.Extension = v
	return s
}

func (s *CreateMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
