// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMessageGroupRequest
	GetAppId() *string
	SetExtension(v map[string]*string) *UpdateMessageGroupRequest
	GetExtension() map[string]*string
	SetGroupId(v string) *UpdateMessageGroupRequest
	GetGroupId() *string
}

type UpdateMessageGroupRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The extended field.
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UpdateMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMessageGroupRequest) GetExtension() map[string]*string {
	return s.Extension
}

func (s *UpdateMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateMessageGroupRequest) SetAppId(v string) *UpdateMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMessageGroupRequest) SetExtension(v map[string]*string) *UpdateMessageGroupRequest {
	s.Extension = v
	return s
}

func (s *UpdateMessageGroupRequest) SetGroupId(v string) *UpdateMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
