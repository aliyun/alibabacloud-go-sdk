// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOutputFileToResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SaveOutputFileToResourceShrinkRequest
	GetDirectoryId() *string
	SetItemIdsShrink(v string) *SaveOutputFileToResourceShrinkRequest
	GetItemIdsShrink() *string
	SetMode(v string) *SaveOutputFileToResourceShrinkRequest
	GetMode() *string
	SetTenantId(v string) *SaveOutputFileToResourceShrinkRequest
	GetTenantId() *string
}

type SaveOutputFileToResourceShrinkRequest struct {
	// The ID of the target personal folder. If not specified, the user\\"s default folder is automatically resolved.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// itemIds
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ItemIdsShrink *string `json:"itemIds,omitempty" xml:"itemIds,omitempty"`
	// The save mode. Valid values:
	//
	// - link: Links the resource to the output in a 1:1 idempotent manner. Edits to the output are synchronized to the resource.
	//
	// - copy: Creates a snapshot copy with no limit on the number of copies.
	//
	// example:
	//
	// link
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// string_value
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SaveOutputFileToResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveOutputFileToResourceShrinkRequest) GetItemIdsShrink() *string {
	return s.ItemIdsShrink
}

func (s *SaveOutputFileToResourceShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveOutputFileToResourceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveOutputFileToResourceShrinkRequest) SetDirectoryId(v string) *SaveOutputFileToResourceShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetItemIdsShrink(v string) *SaveOutputFileToResourceShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetMode(v string) *SaveOutputFileToResourceShrinkRequest {
	s.Mode = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetTenantId(v string) *SaveOutputFileToResourceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
