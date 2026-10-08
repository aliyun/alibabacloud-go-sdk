// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceTableMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *GetSourceTableMetaShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *GetSourceTableMetaShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetSourceTableMetaShrinkRequest
	GetOpUserId() *string
	SetQueryShrink(v string) *GetSourceTableMetaShrinkRequest
	GetQueryShrink() *string
}

type GetSourceTableMetaShrinkRequest struct {
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
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
	// This parameter is required.
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s GetSourceTableMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceTableMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSourceTableMetaShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *GetSourceTableMetaShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetSourceTableMetaShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetSourceTableMetaShrinkRequest) GetQueryShrink() *string {
	return s.QueryShrink
}

func (s *GetSourceTableMetaShrinkRequest) SetContextShrink(v string) *GetSourceTableMetaShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *GetSourceTableMetaShrinkRequest) SetOpTenantId(v int64) *GetSourceTableMetaShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSourceTableMetaShrinkRequest) SetOpUserId(v string) *GetSourceTableMetaShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *GetSourceTableMetaShrinkRequest) SetQueryShrink(v string) *GetSourceTableMetaShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *GetSourceTableMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
