// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDomainsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataDomainsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataDomainsShrinkRequest
	GetOpTenantId() *int64
}

type ListDataDomainsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDataDomainsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataDomainsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataDomainsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataDomainsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataDomainsShrinkRequest) SetListQueryShrink(v string) *ListDataDomainsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataDomainsShrinkRequest) SetOpTenantId(v int64) *ListDataDomainsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataDomainsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
