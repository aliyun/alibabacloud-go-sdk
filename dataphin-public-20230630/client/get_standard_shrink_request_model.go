// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetStandardShrinkRequest
	GetOpTenantId() *int64
	SetStandardGetQueryShrink(v string) *GetStandardShrinkRequest
	GetStandardGetQueryShrink() *string
}

type GetStandardShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	StandardGetQueryShrink *string `json:"StandardGetQuery,omitempty" xml:"StandardGetQuery,omitempty"`
}

func (s GetStandardShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStandardShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetStandardShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStandardShrinkRequest) GetStandardGetQueryShrink() *string {
	return s.StandardGetQueryShrink
}

func (s *GetStandardShrinkRequest) SetOpTenantId(v int64) *GetStandardShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStandardShrinkRequest) SetStandardGetQueryShrink(v string) *GetStandardShrinkRequest {
	s.StandardGetQueryShrink = &v
	return s
}

func (s *GetStandardShrinkRequest) Validate() error {
	return dara.Validate(s)
}
