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
	// 目标个人目录 ID；不传则自动解析用户默认目录。
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
	// 保存方式：link=链接（1:1 幂等，编辑产出会同步资源） / copy=复制（不限次，快照）
	//
	// example:
	//
	// link
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
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
