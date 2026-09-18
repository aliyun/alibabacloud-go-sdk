// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveGroupResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *MoveGroupResourceRequest
	GetGroupId() *string
	SetSourceDirectoryId(v string) *MoveGroupResourceRequest
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveGroupResourceRequest
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveGroupResourceRequest
	GetTargetDirectoryId() *string
	SetTenantId(v string) *MoveGroupResourceRequest
	GetTenantId() *string
}

type MoveGroupResourceRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 资料当前所在的空间物理目录真实 ID，不支持 root 哨兵
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// 待移动的物理 GROUP 资料 ID；引用资料只读
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 同一空间目标物理目录真实 ID，必须与源目录不同
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MoveGroupResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveGroupResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveGroupResourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *MoveGroupResourceRequest) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveGroupResourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveGroupResourceRequest) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveGroupResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MoveGroupResourceRequest) SetGroupId(v string) *MoveGroupResourceRequest {
	s.GroupId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetSourceDirectoryId(v string) *MoveGroupResourceRequest {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetSourceId(v string) *MoveGroupResourceRequest {
	s.SourceId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetTargetDirectoryId(v string) *MoveGroupResourceRequest {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetTenantId(v string) *MoveGroupResourceRequest {
	s.TenantId = &v
	return s
}

func (s *MoveGroupResourceRequest) Validate() error {
	return dara.Validate(s)
}
