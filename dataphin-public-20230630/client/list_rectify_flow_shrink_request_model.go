// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRectifyFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListRectifyFlowShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListRectifyFlowShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListRectifyFlowShrinkRequest
	GetOpUserId() *string
}

type ListRectifyFlowShrinkRequest struct {
	// The business request parameters.
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
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListRectifyFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRectifyFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRectifyFlowShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListRectifyFlowShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListRectifyFlowShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListRectifyFlowShrinkRequest) SetListQueryShrink(v string) *ListRectifyFlowShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListRectifyFlowShrinkRequest) SetOpTenantId(v int64) *ListRectifyFlowShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListRectifyFlowShrinkRequest) SetOpUserId(v string) *ListRectifyFlowShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListRectifyFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
