// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizEntitiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListBizEntitiesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListBizEntitiesShrinkRequest
	GetOpTenantId() *int64
}

type ListBizEntitiesShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListBizEntitiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBizEntitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListBizEntitiesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListBizEntitiesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListBizEntitiesShrinkRequest) SetListQueryShrink(v string) *ListBizEntitiesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListBizEntitiesShrinkRequest) SetOpTenantId(v int64) *ListBizEntitiesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListBizEntitiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
