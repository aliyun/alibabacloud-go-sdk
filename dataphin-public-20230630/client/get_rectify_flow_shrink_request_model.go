// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRectifyFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetQueryShrink(v string) *GetRectifyFlowShrinkRequest
	GetGetQueryShrink() *string
	SetOpTenantId(v int64) *GetRectifyFlowShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetRectifyFlowShrinkRequest
	GetOpUserId() *string
}

type GetRectifyFlowShrinkRequest struct {
	// The business request parameters.
	//
	// This parameter is required.
	GetQueryShrink *string `json:"GetQuery,omitempty" xml:"GetQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetRectifyFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowShrinkRequest) GetGetQueryShrink() *string {
	return s.GetQueryShrink
}

func (s *GetRectifyFlowShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetRectifyFlowShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetRectifyFlowShrinkRequest) SetGetQueryShrink(v string) *GetRectifyFlowShrinkRequest {
	s.GetQueryShrink = &v
	return s
}

func (s *GetRectifyFlowShrinkRequest) SetOpTenantId(v int64) *GetRectifyFlowShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetRectifyFlowShrinkRequest) SetOpUserId(v string) *GetRectifyFlowShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *GetRectifyFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
