// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetOnlineUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *BatchGetOnlineUsersRequest
	GetAppId() *string
	SetGroupId(v string) *BatchGetOnlineUsersRequest
	GetGroupId() *string
	SetUserIds(v string) *BatchGetOnlineUsersRequest
	GetUserIds() *string
}

type BatchGetOnlineUsersRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23wcaec-***695ef
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of user IDs. Separate multiple user IDs with commas (,). You can specify a maximum of 20 user IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// de1**a0,hu**9
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s BatchGetOnlineUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetOnlineUsersRequest) GoString() string {
	return s.String()
}

func (s *BatchGetOnlineUsersRequest) GetAppId() *string {
	return s.AppId
}

func (s *BatchGetOnlineUsersRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchGetOnlineUsersRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *BatchGetOnlineUsersRequest) SetAppId(v string) *BatchGetOnlineUsersRequest {
	s.AppId = &v
	return s
}

func (s *BatchGetOnlineUsersRequest) SetGroupId(v string) *BatchGetOnlineUsersRequest {
	s.GroupId = &v
	return s
}

func (s *BatchGetOnlineUsersRequest) SetUserIds(v string) *BatchGetOnlineUsersRequest {
	s.UserIds = &v
	return s
}

func (s *BatchGetOnlineUsersRequest) Validate() error {
	return dara.Validate(s)
}
