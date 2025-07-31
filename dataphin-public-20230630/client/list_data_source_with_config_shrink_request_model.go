// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceWithConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataSourceWithConfigShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataSourceWithConfigShrinkRequest
	GetOpTenantId() *int64
}

type ListDataSourceWithConfigShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDataSourceWithConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataSourceWithConfigShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataSourceWithConfigShrinkRequest) SetListQueryShrink(v string) *ListDataSourceWithConfigShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataSourceWithConfigShrinkRequest) SetOpTenantId(v int64) *ListDataSourceWithConfigShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataSourceWithConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
