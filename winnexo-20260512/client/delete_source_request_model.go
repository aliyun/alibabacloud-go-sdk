// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *DeleteSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *DeleteSourceRequest
	GetTenantId() *string
}

type DeleteSourceRequest struct {
	// 待删除的数据源 ID（租户内唯一）
	//
	// This parameter is required.
	//
	// example:
	//
	// 781
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 21577
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *DeleteSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteSourceRequest) SetSourceId(v string) *DeleteSourceRequest {
	s.SourceId = &v
	return s
}

func (s *DeleteSourceRequest) SetTenantId(v string) *DeleteSourceRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteSourceRequest) Validate() error {
	return dara.Validate(s)
}
