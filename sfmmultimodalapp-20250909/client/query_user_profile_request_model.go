// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryUserProfileRequest
	GetAppId() *string
	SetUserDefinedId(v string) *QueryUserProfileRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *QueryUserProfileRequest
	GetWorkspaceId() *string
}

type QueryUserProfileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f7e110b6d435
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryUserProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserProfileRequest) GoString() string {
	return s.String()
}

func (s *QueryUserProfileRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryUserProfileRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *QueryUserProfileRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryUserProfileRequest) SetAppId(v string) *QueryUserProfileRequest {
	s.AppId = &v
	return s
}

func (s *QueryUserProfileRequest) SetUserDefinedId(v string) *QueryUserProfileRequest {
	s.UserDefinedId = &v
	return s
}

func (s *QueryUserProfileRequest) SetWorkspaceId(v string) *QueryUserProfileRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryUserProfileRequest) Validate() error {
	return dara.Validate(s)
}
