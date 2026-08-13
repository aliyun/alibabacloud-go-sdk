// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceSync(v bool) *ReparseSourceRequest
	GetForceSync() *bool
	SetSourceId(v string) *ReparseSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *ReparseSourceRequest
	GetTenantId() *string
}

type ReparseSourceRequest struct {
	// 是否同步等待重新解析完成；默认 false，异步入队
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// 待重新解析的数据源 ID（租户内唯一）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReparseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReparseSourceRequest) GoString() string {
	return s.String()
}

func (s *ReparseSourceRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReparseSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReparseSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReparseSourceRequest) SetForceSync(v bool) *ReparseSourceRequest {
	s.ForceSync = &v
	return s
}

func (s *ReparseSourceRequest) SetSourceId(v string) *ReparseSourceRequest {
	s.SourceId = &v
	return s
}

func (s *ReparseSourceRequest) SetTenantId(v string) *ReparseSourceRequest {
	s.TenantId = &v
	return s
}

func (s *ReparseSourceRequest) Validate() error {
	return dara.Validate(s)
}
