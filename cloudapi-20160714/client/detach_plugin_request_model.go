// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DetachPluginRequest
	GetApiId() *string
	SetGroupId(v string) *DetachPluginRequest
	GetGroupId() *string
	SetPluginId(v string) *DetachPluginRequest
	GetPluginId() *string
	SetSecurityToken(v string) *DetachPluginRequest
	GetSecurityToken() *string
	SetStageName(v string) *DetachPluginRequest
	GetStageName() *string
}

type DetachPluginRequest struct {
	// example:
	//
	// 19a2846d8e8541c788c6be740035eb68
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// 93b87279e54c4c0baeb7113cdf9c67f5
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 05df2b52a3644a3a8b1935ab8ab59e9d
	PluginId      *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DetachPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachPluginRequest) GoString() string {
	return s.String()
}

func (s *DetachPluginRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DetachPluginRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DetachPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DetachPluginRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DetachPluginRequest) GetStageName() *string {
	return s.StageName
}

func (s *DetachPluginRequest) SetApiId(v string) *DetachPluginRequest {
	s.ApiId = &v
	return s
}

func (s *DetachPluginRequest) SetGroupId(v string) *DetachPluginRequest {
	s.GroupId = &v
	return s
}

func (s *DetachPluginRequest) SetPluginId(v string) *DetachPluginRequest {
	s.PluginId = &v
	return s
}

func (s *DetachPluginRequest) SetSecurityToken(v string) *DetachPluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *DetachPluginRequest) SetStageName(v string) *DetachPluginRequest {
	s.StageName = &v
	return s
}

func (s *DetachPluginRequest) Validate() error {
	return dara.Validate(s)
}
