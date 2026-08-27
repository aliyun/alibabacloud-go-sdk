// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOutputFileToResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SaveOutputFileToResourceRequest
	GetDirectoryId() *string
	SetItemIds(v []*string) *SaveOutputFileToResourceRequest
	GetItemIds() []*string
	SetMode(v string) *SaveOutputFileToResourceRequest
	GetMode() *string
	SetTenantId(v string) *SaveOutputFileToResourceRequest
	GetTenantId() *string
}

type SaveOutputFileToResourceRequest struct {
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
	ItemIds []*string `json:"itemIds,omitempty" xml:"itemIds,omitempty" type:"Repeated"`
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

func (s SaveOutputFileToResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceRequest) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveOutputFileToResourceRequest) GetItemIds() []*string {
	return s.ItemIds
}

func (s *SaveOutputFileToResourceRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveOutputFileToResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveOutputFileToResourceRequest) SetDirectoryId(v string) *SaveOutputFileToResourceRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetItemIds(v []*string) *SaveOutputFileToResourceRequest {
	s.ItemIds = v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetMode(v string) *SaveOutputFileToResourceRequest {
	s.Mode = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetTenantId(v string) *SaveOutputFileToResourceRequest {
	s.TenantId = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) Validate() error {
	return dara.Validate(s)
}
