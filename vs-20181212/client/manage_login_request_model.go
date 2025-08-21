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
	// example:
	//
	// open
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// example:
	//
	// mygroup
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
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
