// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGroupPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DetachGroupPluginRequest
	GetGroupId() *string
	SetPluginId(v string) *DetachGroupPluginRequest
	GetPluginId() *string
	SetSecurityToken(v string) *DetachGroupPluginRequest
	GetSecurityToken() *string
	SetStageName(v string) *DetachGroupPluginRequest
	GetStageName() *string
}

type DetachGroupPluginRequest struct {
	// API group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// API Gateway plugin ID
	//
	// This parameter is required.
	//
	// example:
	//
	// a96926e82f994915a8da40a119374537
	PluginId      *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specify the environment of the API to operate on.
	//
	// - **RELEASE**: Production
	//
	// - **PRE**: Pre-release
	//
	// - **TEST**: Test
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DetachGroupPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachGroupPluginRequest) GoString() string {
	return s.String()
}

func (s *DetachGroupPluginRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DetachGroupPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DetachGroupPluginRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DetachGroupPluginRequest) GetStageName() *string {
	return s.StageName
}

func (s *DetachGroupPluginRequest) SetGroupId(v string) *DetachGroupPluginRequest {
	s.GroupId = &v
	return s
}

func (s *DetachGroupPluginRequest) SetPluginId(v string) *DetachGroupPluginRequest {
	s.PluginId = &v
	return s
}

func (s *DetachGroupPluginRequest) SetSecurityToken(v string) *DetachGroupPluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *DetachGroupPluginRequest) SetStageName(v string) *DetachGroupPluginRequest {
	s.StageName = &v
	return s
}

func (s *DetachGroupPluginRequest) Validate() error {
	return dara.Validate(s)
}
