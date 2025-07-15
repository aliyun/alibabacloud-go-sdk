// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupUserByIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMessageGroupUserByIdShrinkRequest
	GetAppId() *string
	SetGroupId(v string) *ListMessageGroupUserByIdShrinkRequest
	GetGroupId() *string
	SetUserIdListShrink(v string) *ListMessageGroupUserByIdShrinkRequest
	GetUserIdListShrink() *string
}

type ListMessageGroupUserByIdShrinkRequest struct {
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
	// The list of users.
	//
	// This parameter is required.
	UserIdListShrink *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s ListMessageGroupUserByIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserByIdShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMessageGroupUserByIdShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListMessageGroupUserByIdShrinkRequest) GetUserIdListShrink() *string {
	return s.UserIdListShrink
}

func (s *ListMessageGroupUserByIdShrinkRequest) SetAppId(v string) *ListMessageGroupUserByIdShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListMessageGroupUserByIdShrinkRequest) SetGroupId(v string) *ListMessageGroupUserByIdShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListMessageGroupUserByIdShrinkRequest) SetUserIdListShrink(v string) *ListMessageGroupUserByIdShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

func (s *ListMessageGroupUserByIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
