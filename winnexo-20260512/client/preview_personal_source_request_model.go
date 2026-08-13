// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPersonalSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *PreviewPersonalSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *PreviewPersonalSourceRequest
	GetTenantId() *string
}

type PreviewPersonalSourceRequest struct {
	// 知识 ID（数据源唯一标识）
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

func (s PreviewPersonalSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewPersonalSourceRequest) GoString() string {
	return s.String()
}

func (s *PreviewPersonalSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *PreviewPersonalSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewPersonalSourceRequest) SetSourceId(v string) *PreviewPersonalSourceRequest {
	s.SourceId = &v
	return s
}

func (s *PreviewPersonalSourceRequest) SetTenantId(v string) *PreviewPersonalSourceRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewPersonalSourceRequest) Validate() error {
	return dara.Validate(s)
}
