// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryProfileRequest
	GetAppId() *string
	SetUserDefinedId(v string) *QueryProfileRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *QueryProfileRequest
	GetWorkspaceId() *string
}

type QueryProfileRequest struct {
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
	// 6d4359977d1
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryProfileRequest) GoString() string {
	return s.String()
}

func (s *QueryProfileRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryProfileRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *QueryProfileRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryProfileRequest) SetAppId(v string) *QueryProfileRequest {
	s.AppId = &v
	return s
}

func (s *QueryProfileRequest) SetUserDefinedId(v string) *QueryProfileRequest {
	s.UserDefinedId = &v
	return s
}

func (s *QueryProfileRequest) SetWorkspaceId(v string) *QueryProfileRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryProfileRequest) Validate() error {
	return dara.Validate(s)
}
