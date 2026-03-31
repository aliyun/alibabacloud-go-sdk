// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMemoryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMemoryConfigRequest
	GetAppId() *string
	SetUserDefinedId(v string) *QueryMemoryConfigRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *QueryMemoryConfigRequest
	GetWorkspaceId() *string
}

type QueryMemoryConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 110b6d4359977d1
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMemoryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryMemoryConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMemoryConfigRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *QueryMemoryConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMemoryConfigRequest) SetAppId(v string) *QueryMemoryConfigRequest {
	s.AppId = &v
	return s
}

func (s *QueryMemoryConfigRequest) SetUserDefinedId(v string) *QueryMemoryConfigRequest {
	s.UserDefinedId = &v
	return s
}

func (s *QueryMemoryConfigRequest) SetWorkspaceId(v string) *QueryMemoryConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMemoryConfigRequest) Validate() error {
	return dara.Validate(s)
}
