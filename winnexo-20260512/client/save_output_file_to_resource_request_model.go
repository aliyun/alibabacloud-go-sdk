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
	ItemIds []*string `json:"itemIds,omitempty" xml:"itemIds,omitempty" type:"Repeated"`
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
