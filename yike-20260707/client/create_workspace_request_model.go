// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTitle(v string) *CreateWorkspaceRequest
	GetTitle() *string
	SetUserCountLimit(v int64) *CreateWorkspaceRequest
	GetUserCountLimit() *int64
}

type CreateWorkspaceRequest struct {
	// The name of the studio. The name can be up to 50 characters in length and cannot contain single quotation marks, semicolons, less than signs, greater than signs, equal signs, vertical lines, or parentheses.
	//
	// This parameter is required.
	//
	// example:
	//
	// Video production studio
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The maximum number of users allowed in the studio team. Valid values: 2 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	UserCountLimit *int64 `json:"UserCountLimit,omitempty" xml:"UserCountLimit,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkspaceRequest) GetUserCountLimit() *int64 {
	return s.UserCountLimit
}

func (s *CreateWorkspaceRequest) SetTitle(v string) *CreateWorkspaceRequest {
	s.Title = &v
	return s
}

func (s *CreateWorkspaceRequest) SetUserCountLimit(v int64) *CreateWorkspaceRequest {
	s.UserCountLimit = &v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
