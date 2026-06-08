// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSecurityStrategyShrinkRequest
	GetClientToken() *string
	SetContentShrink(v string) *UpdateSecurityStrategyShrinkRequest
	GetContentShrink() *string
	SetDescription(v string) *UpdateSecurityStrategyShrinkRequest
	GetDescription() *string
	SetId(v int64) *UpdateSecurityStrategyShrinkRequest
	GetId() *int64
	SetName(v string) *UpdateSecurityStrategyShrinkRequest
	GetName() *string
	SetWorkspacesShrink(v string) *UpdateSecurityStrategyShrinkRequest
	GetWorkspacesShrink() *string
}

type UpdateSecurityStrategyShrinkRequest struct {
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkspacesShrink *string `json:"Workspaces,omitempty" xml:"Workspaces,omitempty"`
}

func (s UpdateSecurityStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityStrategyShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSecurityStrategyShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *UpdateSecurityStrategyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecurityStrategyShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateSecurityStrategyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSecurityStrategyShrinkRequest) GetWorkspacesShrink() *string {
	return s.WorkspacesShrink
}

func (s *UpdateSecurityStrategyShrinkRequest) SetClientToken(v string) *UpdateSecurityStrategyShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSecurityStrategyShrinkRequest) SetContentShrink(v string) *UpdateSecurityStrategyShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *UpdateSecurityStrategyShrinkRequest) SetDescription(v string) *UpdateSecurityStrategyShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecurityStrategyShrinkRequest) SetId(v int64) *UpdateSecurityStrategyShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateSecurityStrategyShrinkRequest) SetName(v string) *UpdateSecurityStrategyShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSecurityStrategyShrinkRequest) SetWorkspacesShrink(v string) *UpdateSecurityStrategyShrinkRequest {
	s.WorkspacesShrink = &v
	return s
}

func (s *UpdateSecurityStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
