// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateLinkRequest
	GetAppId() *string
	SetCors(v bool) *UpdateLinkRequest
	GetCors() *bool
	SetDomain(v string) *UpdateLinkRequest
	GetDomain() *string
	SetDynamicfield(v string) *UpdateLinkRequest
	GetDynamicfield() *string
	SetTargetUrl(v string) *UpdateLinkRequest
	GetTargetUrl() *string
	SetUrl(v string) *UpdateLinkRequest
	GetUrl() *string
	SetWorkspaceId(v string) *UpdateLinkRequest
	GetWorkspaceId() *string
}

type UpdateLinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BB5953C300957
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// true
	Cors *bool `json:"Cors,omitempty" xml:"Cors,omitempty"`
	// example:
	//
	// x519.cn
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// txt
	Dynamicfield *string `json:"Dynamicfield,omitempty" xml:"Dynamicfield,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://********
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx/8hGb9SyJARqp7V4PGP92X
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateLinkRequest) GetCors() *bool {
	return s.Cors
}

func (s *UpdateLinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateLinkRequest) GetDynamicfield() *string {
	return s.Dynamicfield
}

func (s *UpdateLinkRequest) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateLinkRequest) GetUrl() *string {
	return s.Url
}

func (s *UpdateLinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateLinkRequest) SetAppId(v string) *UpdateLinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLinkRequest) SetCors(v bool) *UpdateLinkRequest {
	s.Cors = &v
	return s
}

func (s *UpdateLinkRequest) SetDomain(v string) *UpdateLinkRequest {
	s.Domain = &v
	return s
}

func (s *UpdateLinkRequest) SetDynamicfield(v string) *UpdateLinkRequest {
	s.Dynamicfield = &v
	return s
}

func (s *UpdateLinkRequest) SetTargetUrl(v string) *UpdateLinkRequest {
	s.TargetUrl = &v
	return s
}

func (s *UpdateLinkRequest) SetUrl(v string) *UpdateLinkRequest {
	s.Url = &v
	return s
}

func (s *UpdateLinkRequest) SetWorkspaceId(v string) *UpdateLinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateLinkRequest) Validate() error {
	return dara.Validate(s)
}
