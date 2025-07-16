// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFolderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddFolderShrinkRequest
	GetName() *string
	SetOptionShrink(v string) *AddFolderShrinkRequest
	GetOptionShrink() *string
	SetParentId(v string) *AddFolderShrinkRequest
	GetParentId() *string
	SetSpaceId(v string) *AddFolderShrinkRequest
	GetSpaceId() *string
	SetTenantContextShrink(v string) *AddFolderShrinkRequest
	GetTenantContextShrink() *string
}

type AddFolderShrinkRequest struct {
	// This parameter is required.
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OptionShrink *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 140822073803
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xPar2SZ63KodG3aV
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s AddFolderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFolderShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddFolderShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddFolderShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *AddFolderShrinkRequest) GetParentId() *string {
	return s.ParentId
}

func (s *AddFolderShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *AddFolderShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddFolderShrinkRequest) SetName(v string) *AddFolderShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddFolderShrinkRequest) SetOptionShrink(v string) *AddFolderShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *AddFolderShrinkRequest) SetParentId(v string) *AddFolderShrinkRequest {
	s.ParentId = &v
	return s
}

func (s *AddFolderShrinkRequest) SetSpaceId(v string) *AddFolderShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *AddFolderShrinkRequest) SetTenantContextShrink(v string) *AddFolderShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddFolderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
