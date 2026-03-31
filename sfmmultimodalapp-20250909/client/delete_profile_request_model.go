// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteProfileRequest
	GetAppId() *string
	SetUserDefinedId(v string) *DeleteProfileRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *DeleteProfileRequest
	GetWorkspaceId() *string
}

type DeleteProfileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 5A7C969B-5101-112B-8202-DFAEEA4BFBED
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteProfileRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteProfileRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *DeleteProfileRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteProfileRequest) SetAppId(v string) *DeleteProfileRequest {
	s.AppId = &v
	return s
}

func (s *DeleteProfileRequest) SetUserDefinedId(v string) *DeleteProfileRequest {
	s.UserDefinedId = &v
	return s
}

func (s *DeleteProfileRequest) SetWorkspaceId(v string) *DeleteProfileRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteProfileRequest) Validate() error {
	return dara.Validate(s)
}
