// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToPersonalResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetDirectoryId() *string
	SetGroupId(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetGroupId() *string
	SetItemIds(v []*string) *SaveGroupOutputFileToPersonalResourceRequest
	GetItemIds() []*string
	SetMode(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetMode() *string
	SetTenantId(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetTenantId() *string
}

type SaveGroupOutputFileToPersonalResourceRequest struct {
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
	ItemIds []*string `json:"itemIds,omitempty" xml:"itemIds,omitempty" type:"Repeated"`
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

func (s SaveGroupOutputFileToPersonalResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToPersonalResourceRequest) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetItemIds() []*string {
	return s.ItemIds
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetDirectoryId(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetGroupId(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.GroupId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetItemIds(v []*string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.ItemIds = v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetMode(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.Mode = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetTenantId(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) Validate() error {
	return dara.Validate(s)
}
