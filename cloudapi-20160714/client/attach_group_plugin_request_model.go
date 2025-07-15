// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachGroupPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *AttachGroupPluginRequest
	GetGroupId() *string
	SetPluginId(v string) *AttachGroupPluginRequest
	GetPluginId() *string
	SetSecurityToken(v string) *AttachGroupPluginRequest
	GetSecurityToken() *string
	SetStageName(v string) *AttachGroupPluginRequest
	GetStageName() *string
}

type AttachGroupPluginRequest struct {
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 08ae4aa0f95e4321849ee57f4e0b3077
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the plug-in to be bound.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05df2b52a3644a3a8b1935ab8ab59e9d
	PluginId      *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment in which the API is requested. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the pre-release environment
	//
	// 	- **TEST**: the test environment
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s AttachGroupPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachGroupPluginRequest) GoString() string {
	return s.String()
}

func (s *AttachGroupPluginRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AttachGroupPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *AttachGroupPluginRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AttachGroupPluginRequest) GetStageName() *string {
	return s.StageName
}

func (s *AttachGroupPluginRequest) SetGroupId(v string) *AttachGroupPluginRequest {
	s.GroupId = &v
	return s
}

func (s *AttachGroupPluginRequest) SetPluginId(v string) *AttachGroupPluginRequest {
	s.PluginId = &v
	return s
}

func (s *AttachGroupPluginRequest) SetSecurityToken(v string) *AttachGroupPluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *AttachGroupPluginRequest) SetStageName(v string) *AttachGroupPluginRequest {
	s.StageName = &v
	return s
}

func (s *AttachGroupPluginRequest) Validate() error {
	return dara.Validate(s)
}
