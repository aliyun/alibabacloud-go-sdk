// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageLoginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionName(v string) *ManageLoginRequest
	GetActionName() *string
	SetKeyGroup(v string) *ManageLoginRequest
	GetKeyGroup() *string
	SetKeyName(v string) *ManageLoginRequest
	GetKeyName() *string
	SetRenderingInstanceId(v string) *ManageLoginRequest
	GetRenderingInstanceId() *string
}

type ManageLoginRequest struct {
	// Name of the management action. Valid values:
	//
	// 1. open — Activate the public key. This is the default value.
	//
	// 2. close — Deactivate the public key.
	//
	// example:
	//
	// open
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// Name of the public key group. If you do not specify KeyName, all public keys in this group are applied.
	//
	// example:
	//
	// mygroup
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// Name of the public key. You must specify either KeyName or KeyGroup.
	//
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// ID of the Cloud Application Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s ManageLoginRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageLoginRequest) GoString() string {
	return s.String()
}

func (s *ManageLoginRequest) GetActionName() *string {
	return s.ActionName
}

func (s *ManageLoginRequest) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *ManageLoginRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *ManageLoginRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ManageLoginRequest) SetActionName(v string) *ManageLoginRequest {
	s.ActionName = &v
	return s
}

func (s *ManageLoginRequest) SetKeyGroup(v string) *ManageLoginRequest {
	s.KeyGroup = &v
	return s
}

func (s *ManageLoginRequest) SetKeyName(v string) *ManageLoginRequest {
	s.KeyName = &v
	return s
}

func (s *ManageLoginRequest) SetRenderingInstanceId(v string) *ManageLoginRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ManageLoginRequest) Validate() error {
	return dara.Validate(s)
}
