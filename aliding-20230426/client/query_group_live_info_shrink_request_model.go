// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupLiveInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnchorUnionId(v string) *QueryGroupLiveInfoShrinkRequest
	GetAnchorUnionId() *string
	SetLiveUuid(v string) *QueryGroupLiveInfoShrinkRequest
	GetLiveUuid() *string
	SetTenantContextShrink(v string) *QueryGroupLiveInfoShrinkRequest
	GetTenantContextShrink() *string
}

type QueryGroupLiveInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Eijxxx
	AnchorUnionId *string `json:"AnchorUnionId,omitempty" xml:"AnchorUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	LiveUuid            *string `json:"LiveUuid,omitempty" xml:"LiveUuid,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryGroupLiveInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupLiveInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveInfoShrinkRequest) GetAnchorUnionId() *string {
	return s.AnchorUnionId
}

func (s *QueryGroupLiveInfoShrinkRequest) GetLiveUuid() *string {
	return s.LiveUuid
}

func (s *QueryGroupLiveInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryGroupLiveInfoShrinkRequest) SetAnchorUnionId(v string) *QueryGroupLiveInfoShrinkRequest {
	s.AnchorUnionId = &v
	return s
}

func (s *QueryGroupLiveInfoShrinkRequest) SetLiveUuid(v string) *QueryGroupLiveInfoShrinkRequest {
	s.LiveUuid = &v
	return s
}

func (s *QueryGroupLiveInfoShrinkRequest) SetTenantContextShrink(v string) *QueryGroupLiveInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryGroupLiveInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
