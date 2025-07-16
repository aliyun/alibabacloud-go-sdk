// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLatestPlanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetUserLatestPlanShrinkRequest
	GetTenantContextShrink() *string
}

type GetUserLatestPlanShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetUserLatestPlanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserLatestPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserLatestPlanShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetUserLatestPlanShrinkRequest) SetTenantContextShrink(v string) *GetUserLatestPlanShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetUserLatestPlanShrinkRequest) Validate() error {
	return dara.Validate(s)
}
