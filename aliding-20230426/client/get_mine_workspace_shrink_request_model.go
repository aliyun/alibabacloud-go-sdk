// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMineWorkspaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestShrink(v string) *GetMineWorkspaceShrinkRequest
	GetRequestShrink() *string
	SetTenantContextShrink(v string) *GetMineWorkspaceShrinkRequest
	GetTenantContextShrink() *string
}

type GetMineWorkspaceShrinkRequest struct {
	RequestShrink       *string `json:"Request,omitempty" xml:"Request,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetMineWorkspaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceShrinkRequest) GetRequestShrink() *string {
	return s.RequestShrink
}

func (s *GetMineWorkspaceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetMineWorkspaceShrinkRequest) SetRequestShrink(v string) *GetMineWorkspaceShrinkRequest {
	s.RequestShrink = &v
	return s
}

func (s *GetMineWorkspaceShrinkRequest) SetTenantContextShrink(v string) *GetMineWorkspaceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetMineWorkspaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
