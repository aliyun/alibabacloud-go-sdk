// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAssetsOffShelveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitAssetsOffShelveShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitAssetsOffShelveShrinkRequest
	GetOpUserId() *string
	SetSubmitCommandShrink(v string) *SubmitAssetsOffShelveShrinkRequest
	GetSubmitCommandShrink() *string
}

type SubmitAssetsOffShelveShrinkRequest struct {
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
	// The delisting submit command.
	//
	// This parameter is required.
	SubmitCommandShrink *string `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty"`
}

func (s SubmitAssetsOffShelveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOffShelveShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOffShelveShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitAssetsOffShelveShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitAssetsOffShelveShrinkRequest) GetSubmitCommandShrink() *string {
	return s.SubmitCommandShrink
}

func (s *SubmitAssetsOffShelveShrinkRequest) SetOpTenantId(v int64) *SubmitAssetsOffShelveShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitAssetsOffShelveShrinkRequest) SetOpUserId(v string) *SubmitAssetsOffShelveShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitAssetsOffShelveShrinkRequest) SetSubmitCommandShrink(v string) *SubmitAssetsOffShelveShrinkRequest {
	s.SubmitCommandShrink = &v
	return s
}

func (s *SubmitAssetsOffShelveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
