// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflinePipelineByAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *OfflinePipelineByAsyncShrinkRequest
	GetContextShrink() *string
	SetOfflineCommandShrink(v string) *OfflinePipelineByAsyncShrinkRequest
	GetOfflineCommandShrink() *string
	SetOpTenantId(v int64) *OfflinePipelineByAsyncShrinkRequest
	GetOpTenantId() *int64
}

type OfflinePipelineByAsyncShrinkRequest struct {
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	OfflineCommandShrink *string `json:"OfflineCommand,omitempty" xml:"OfflineCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s OfflinePipelineByAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineByAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *OfflinePipelineByAsyncShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *OfflinePipelineByAsyncShrinkRequest) GetOfflineCommandShrink() *string {
	return s.OfflineCommandShrink
}

func (s *OfflinePipelineByAsyncShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflinePipelineByAsyncShrinkRequest) SetContextShrink(v string) *OfflinePipelineByAsyncShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *OfflinePipelineByAsyncShrinkRequest) SetOfflineCommandShrink(v string) *OfflinePipelineByAsyncShrinkRequest {
	s.OfflineCommandShrink = &v
	return s
}

func (s *OfflinePipelineByAsyncShrinkRequest) SetOpTenantId(v int64) *OfflinePipelineByAsyncShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflinePipelineByAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
