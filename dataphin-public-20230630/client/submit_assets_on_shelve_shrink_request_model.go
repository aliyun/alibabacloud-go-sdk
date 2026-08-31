// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAssetsOnShelveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitAssetsOnShelveShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitAssetsOnShelveShrinkRequest
	GetOpUserId() *string
	SetSubmitCommandShrink(v string) *SubmitAssetsOnShelveShrinkRequest
	GetSubmitCommandShrink() *string
}

type SubmitAssetsOnShelveShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The submit listing instruction.
	//
	// This parameter is required.
	SubmitCommandShrink *string `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty"`
}

func (s SubmitAssetsOnShelveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOnShelveShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOnShelveShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitAssetsOnShelveShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitAssetsOnShelveShrinkRequest) GetSubmitCommandShrink() *string {
	return s.SubmitCommandShrink
}

func (s *SubmitAssetsOnShelveShrinkRequest) SetOpTenantId(v int64) *SubmitAssetsOnShelveShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitAssetsOnShelveShrinkRequest) SetOpUserId(v string) *SubmitAssetsOnShelveShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitAssetsOnShelveShrinkRequest) SetSubmitCommandShrink(v string) *SubmitAssetsOnShelveShrinkRequest {
	s.SubmitCommandShrink = &v
	return s
}

func (s *SubmitAssetsOnShelveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
