// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTitle(v string) *CreateYikeWorkspaceRequest
	GetTitle() *string
	SetUserCountLimit(v int64) *CreateYikeWorkspaceRequest
	GetUserCountLimit() *int64
}

type CreateYikeWorkspaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	UserCountLimit *int64 `json:"UserCountLimit,omitempty" xml:"UserCountLimit,omitempty"`
}

func (s CreateYikeWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateYikeWorkspaceRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateYikeWorkspaceRequest) GetUserCountLimit() *int64 {
	return s.UserCountLimit
}

func (s *CreateYikeWorkspaceRequest) SetTitle(v string) *CreateYikeWorkspaceRequest {
	s.Title = &v
	return s
}

func (s *CreateYikeWorkspaceRequest) SetUserCountLimit(v int64) *CreateYikeWorkspaceRequest {
	s.UserCountLimit = &v
	return s
}

func (s *CreateYikeWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
