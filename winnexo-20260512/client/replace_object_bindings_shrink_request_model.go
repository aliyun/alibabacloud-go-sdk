// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceObjectBindingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectBindingsShrink(v string) *ReplaceObjectBindingsShrinkRequest
	GetObjectBindingsShrink() *string
	SetSourceId(v string) *ReplaceObjectBindingsShrinkRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceObjectBindingsShrinkRequest
	GetTenantId() *string
}

type ReplaceObjectBindingsShrinkRequest struct {
	// 新的对象绑定列表（全量替换；传空列表表示清空所有绑定）
	//
	// This parameter is required.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
	// 数据源 ID（租户内唯一）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceObjectBindingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *ReplaceObjectBindingsShrinkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceObjectBindingsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceObjectBindingsShrinkRequest) SetObjectBindingsShrink(v string) *ReplaceObjectBindingsShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetSourceId(v string) *ReplaceObjectBindingsShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) SetTenantId(v string) *ReplaceObjectBindingsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceObjectBindingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
