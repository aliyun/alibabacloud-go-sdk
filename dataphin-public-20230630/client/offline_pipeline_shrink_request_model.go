// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflinePipelineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *OfflinePipelineShrinkRequest
	GetContextShrink() *string
	SetOfflineCommandShrink(v string) *OfflinePipelineShrinkRequest
	GetOfflineCommandShrink() *string
	SetOpTenantId(v int64) *OfflinePipelineShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *OfflinePipelineShrinkRequest
	GetOpUserId() *string
}

type OfflinePipelineShrinkRequest struct {
	// The request context information.
	//
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The offline command for the pipeline node.
	//
	// This parameter is required.
	OfflineCommandShrink *string `json:"OfflineCommand,omitempty" xml:"OfflineCommand,omitempty"`
	// The tenant ID.
	//
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
}

func (s OfflinePipelineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineShrinkRequest) GoString() string {
	return s.String()
}

func (s *OfflinePipelineShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *OfflinePipelineShrinkRequest) GetOfflineCommandShrink() *string {
	return s.OfflineCommandShrink
}

func (s *OfflinePipelineShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflinePipelineShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *OfflinePipelineShrinkRequest) SetContextShrink(v string) *OfflinePipelineShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *OfflinePipelineShrinkRequest) SetOfflineCommandShrink(v string) *OfflinePipelineShrinkRequest {
	s.OfflineCommandShrink = &v
	return s
}

func (s *OfflinePipelineShrinkRequest) SetOpTenantId(v int64) *OfflinePipelineShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflinePipelineShrinkRequest) SetOpUserId(v string) *OfflinePipelineShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *OfflinePipelineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
