// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupUserByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMessageGroupUserByIdRequest
	GetAppId() *string
	SetGroupId(v string) *ListMessageGroupUserByIdRequest
	GetGroupId() *string
	SetUserIdList(v []*string) *ListMessageGroupUserByIdRequest
	GetUserIdList() []*string
}

type ListMessageGroupUserByIdRequest struct {
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
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s ListMessageGroupUserByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserByIdRequest) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserByIdRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMessageGroupUserByIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListMessageGroupUserByIdRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *ListMessageGroupUserByIdRequest) SetAppId(v string) *ListMessageGroupUserByIdRequest {
	s.AppId = &v
	return s
}

func (s *ListMessageGroupUserByIdRequest) SetGroupId(v string) *ListMessageGroupUserByIdRequest {
	s.GroupId = &v
	return s
}

func (s *ListMessageGroupUserByIdRequest) SetUserIdList(v []*string) *ListMessageGroupUserByIdRequest {
	s.UserIdList = v
	return s
}

func (s *ListMessageGroupUserByIdRequest) Validate() error {
	return dara.Validate(s)
}
