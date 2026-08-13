// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceDirectoryId(v string) *MoveResourceRequest
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveResourceRequest
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveResourceRequest
	GetTargetDirectoryId() *string
	SetTenantId(v string) *MoveResourceRequest
	GetTenantId() *string
}

type MoveResourceRequest struct {
	// 源目录 ID（资源当前所在的个人目录）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceDirectoryId
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// 待移动的资源 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 目标目录 ID（资源即将移动到的个人目录）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTargetDirectoryId
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MoveResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceRequest) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveResourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveResourceRequest) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MoveResourceRequest) SetSourceDirectoryId(v string) *MoveResourceRequest {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveResourceRequest) SetSourceId(v string) *MoveResourceRequest {
	s.SourceId = &v
	return s
}

func (s *MoveResourceRequest) SetTargetDirectoryId(v string) *MoveResourceRequest {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveResourceRequest) SetTenantId(v string) *MoveResourceRequest {
	s.TenantId = &v
	return s
}

func (s *MoveResourceRequest) Validate() error {
	return dara.Validate(s)
}
