// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryLinkRequest
	GetAppId() *string
	SetUrl(v string) *QueryLinkRequest
	GetUrl() *string
	SetWorkspaceId(v string) *QueryLinkRequest
	GetWorkspaceId() *string
}

type QueryLinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BB5953C300957
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8hGb9SyJARqp7V4PGP92X
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLinkRequest) GoString() string {
	return s.String()
}

func (s *QueryLinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryLinkRequest) GetUrl() *string {
	return s.Url
}

func (s *QueryLinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryLinkRequest) SetAppId(v string) *QueryLinkRequest {
	s.AppId = &v
	return s
}

func (s *QueryLinkRequest) SetUrl(v string) *QueryLinkRequest {
	s.Url = &v
	return s
}

func (s *QueryLinkRequest) SetWorkspaceId(v string) *QueryLinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryLinkRequest) Validate() error {
	return dara.Validate(s)
}
