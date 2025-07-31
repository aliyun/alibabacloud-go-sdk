// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiByAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListApiByAppShrinkRequest
	GetOpTenantId() *int64
	SetPageQueryShrink(v string) *ListApiByAppShrinkRequest
	GetPageQueryShrink() *string
}

type ListApiByAppShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PageQueryShrink *string `json:"PageQuery,omitempty" xml:"PageQuery,omitempty"`
}

func (s ListApiByAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListApiByAppShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListApiByAppShrinkRequest) GetPageQueryShrink() *string {
	return s.PageQueryShrink
}

func (s *ListApiByAppShrinkRequest) SetOpTenantId(v int64) *ListApiByAppShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListApiByAppShrinkRequest) SetPageQueryShrink(v string) *ListApiByAppShrinkRequest {
	s.PageQueryShrink = &v
	return s
}

func (s *ListApiByAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
