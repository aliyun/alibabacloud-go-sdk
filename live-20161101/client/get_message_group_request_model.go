// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMessageGroupRequest
	GetAppId() *string
	SetGroupId(v string) *GetMessageGroupRequest
	GetGroupId() *string
}

type GetMessageGroupRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *GetMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetMessageGroupRequest) SetAppId(v string) *GetMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *GetMessageGroupRequest) SetGroupId(v string) *GetMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
