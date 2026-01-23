// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStandardsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListStandardsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListStandardsShrinkRequest
	GetOpTenantId() *int64
}

type ListStandardsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListStandardsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListStandardsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListStandardsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListStandardsShrinkRequest) SetListQueryShrink(v string) *ListStandardsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListStandardsShrinkRequest) SetOpTenantId(v int64) *ListStandardsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListStandardsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
