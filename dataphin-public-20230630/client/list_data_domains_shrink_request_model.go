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
	SetOpUserId(v string) *ListDataDomainsShrinkRequest
	GetOpUserId() *string
}

type ListDataDomainsShrinkRequest struct {
	// The query request.
	//
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *ListDataDomainsShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListDataDomainsShrinkRequest) SetListQueryShrink(v string) *ListDataDomainsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataDomainsShrinkRequest) SetOpTenantId(v int64) *ListDataDomainsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataDomainsShrinkRequest) SetOpUserId(v string) *ListDataDomainsShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListDataDomainsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
