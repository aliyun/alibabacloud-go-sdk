// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToPersonalResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetGroupId() *string
	SetItemIdsShrink(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetItemIdsShrink() *string
	SetMode(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetMode() *string
	SetTenantId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetTenantId() *string
}

type SaveGroupOutputFileToPersonalResourceShrinkRequest struct {
	// The enterprise knowledge base directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// itemIds
	//
	// This parameter is required.
	//
	// example:
	//
	// ["item-1","item-2"]
	ItemIdsShrink *string `json:"itemIds,omitempty" xml:"itemIds,omitempty"`
	// The save mode. Valid values:
	//
	// - link: creates a link (1:1 idempotent, editing the output synchronizes the resource).
	//
	// - copy: creates a copy (unlimited times, snapshot).
	//
	// This parameter is required.
	//
	// example:
	//
	// copy
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SaveGroupOutputFileToPersonalResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToPersonalResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetItemIdsShrink() *string {
	return s.ItemIdsShrink
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetDirectoryId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetGroupId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetItemIdsShrink(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetMode(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.Mode = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetTenantId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
