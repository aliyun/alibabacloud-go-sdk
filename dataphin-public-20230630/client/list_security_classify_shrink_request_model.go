// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityClassifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListSecurityClassifyShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListSecurityClassifyShrinkRequest
	GetOpTenantId() *int64
}

type ListSecurityClassifyShrinkRequest struct {
	// The query conditions.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListSecurityClassifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityClassifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityClassifyShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListSecurityClassifyShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListSecurityClassifyShrinkRequest) SetListQueryShrink(v string) *ListSecurityClassifyShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListSecurityClassifyShrinkRequest) SetOpTenantId(v int64) *ListSecurityClassifyShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListSecurityClassifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
