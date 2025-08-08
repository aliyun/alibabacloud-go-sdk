// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateLinkRequest
	GetAppId() *string
	SetCors(v bool) *CreateLinkRequest
	GetCors() *bool
	SetDomain(v string) *CreateLinkRequest
	GetDomain() *string
	SetDynamicfield(v string) *CreateLinkRequest
	GetDynamicfield() *string
	SetTargetUrl(v string) *CreateLinkRequest
	GetTargetUrl() *string
	SetWorkspaceId(v string) *CreateLinkRequest
	GetWorkspaceId() *string
}

type CreateLinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BB5953C300957
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// false
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
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateLinkRequest) GetCors() *bool {
	return s.Cors
}

func (s *CreateLinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateLinkRequest) GetDynamicfield() *string {
	return s.Dynamicfield
}

func (s *CreateLinkRequest) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateLinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateLinkRequest) SetAppId(v string) *CreateLinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateLinkRequest) SetCors(v bool) *CreateLinkRequest {
	s.Cors = &v
	return s
}

func (s *CreateLinkRequest) SetDomain(v string) *CreateLinkRequest {
	s.Domain = &v
	return s
}

func (s *CreateLinkRequest) SetDynamicfield(v string) *CreateLinkRequest {
	s.Dynamicfield = &v
	return s
}

func (s *CreateLinkRequest) SetTargetUrl(v string) *CreateLinkRequest {
	s.TargetUrl = &v
	return s
}

func (s *CreateLinkRequest) SetWorkspaceId(v string) *CreateLinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateLinkRequest) Validate() error {
	return dara.Validate(s)
}
