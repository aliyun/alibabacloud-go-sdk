// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseGroupSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceSync(v bool) *ReparseGroupSourceRequest
	GetForceSync() *bool
	SetGroupId(v string) *ReparseGroupSourceRequest
	GetGroupId() *string
	SetSourceId(v string) *ReparseGroupSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *ReparseGroupSourceRequest
	GetTenantId() *string
}

type ReparseGroupSourceRequest struct {
	// 是否等待解析完成；默认 false 异步受理，true 同步等待，网关超时 300000ms
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// 资料所属协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 当前空间物理 GROUP 资料 ID；引用资料只读
	//
	// This parameter is required.
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReparseGroupSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReparseGroupSourceRequest) GoString() string {
	return s.String()
}

func (s *ReparseGroupSourceRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReparseGroupSourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReparseGroupSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReparseGroupSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReparseGroupSourceRequest) SetForceSync(v bool) *ReparseGroupSourceRequest {
	s.ForceSync = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetGroupId(v string) *ReparseGroupSourceRequest {
	s.GroupId = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetSourceId(v string) *ReparseGroupSourceRequest {
	s.SourceId = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetTenantId(v string) *ReparseGroupSourceRequest {
	s.TenantId = &v
	return s
}

func (s *ReparseGroupSourceRequest) Validate() error {
	return dara.Validate(s)
}
