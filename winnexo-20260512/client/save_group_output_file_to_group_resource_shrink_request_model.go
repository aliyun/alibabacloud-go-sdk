// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToGroupResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetGroupId() *string
	SetItemIdsShrink(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetItemIdsShrink() *string
	SetMode(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetMode() *string
	SetTenantId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetTenantId() *string
}

type SaveGroupOutputFileToGroupResourceShrinkRequest struct {
	// The ID of the target personal directory. If not specified, the user\\"s default directory is automatically resolved.
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
	// link
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SaveGroupOutputFileToGroupResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToGroupResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetItemIdsShrink() *string {
	return s.ItemIdsShrink
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetDirectoryId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetGroupId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetItemIdsShrink(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetMode(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.Mode = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetTenantId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
